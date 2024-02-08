package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/gorm"
	"log"
	"math"
	"os"
	"sync"
	"team00/internal/consts"
	"team00/internal/models"
	pb "team00/internal/transmitter/proto"
	"time"
)

const qualityOfValues uint64 = 50 // 50 - 100

type DetectorOfAnomalies struct {
	mu                 sync.Mutex
	Db                 *gorm.DB
	Anomaly            models.Anomaly
	Mean               float64
	STD                float64
	SumOfMean          float64
	SumOfSTD           float64
	anomalyCoefficient float64
	Values             []float64
	qualityOfValues    uint64
}

type Counts struct {
	CountOfRecords   uint64
	CountOfAnomalies uint64
}

func (a *DetectorOfAnomalies) CalculationMeanSTD(frequency float64, count *Counts) {
	a.mu.Lock()
	defer a.mu.Unlock()

	count.CountOfRecords++
	previousMean := a.Mean
	a.SumOfMean += frequency
	a.Mean = a.SumOfMean / float64(count.CountOfRecords)
	a.Values = append(a.Values, frequency)

	diff := frequency - previousMean
	a.SumOfSTD += diff * diff

	a.STD = a.calculateStandardDeviation(a.Values, a.Mean)

	if count.CountOfRecords > a.qualityOfValues {
		if math.Abs(frequency-a.Mean) > a.anomalyCoefficient*a.STD {
			fmt.Printf("Anomaly: %.4f\n", frequency)
			a.Db.Create(a.Anomaly)
			count.CountOfAnomalies++
		}
	}
}

func (a *DetectorOfAnomalies) calculateStandardDeviation(values []float64, mean float64) float64 {
	sumOfSquaredDiffs := 0.0
	for _, value := range values {
		diff := value - mean
		sumOfSquaredDiffs += diff * diff
	}
	variance := sumOfSquaredDiffs / float64(len(values)-1)
	return math.Sqrt(variance)
}

func (a *DetectorOfAnomalies) Initialize(count *Counts) {
	a.mu.Lock()
	defer a.mu.Unlock()

	count.CountOfRecords = 0
	count.CountOfAnomalies = 0
	a.Mean = 0
	a.STD = 0
	a.SumOfMean = 0
	a.SumOfSTD = 0
	a.Values = make([]float64, 0, a.qualityOfValues)
}

func (a *DetectorOfAnomalies) DetectAnomalies(frequencyCh chan float64, sigChan chan bool, count *Counts) {
	fmt.Printf("Calculation of parameters. Enter %d frequencies.\n", a.qualityOfValues)

	for frequency := range frequencyCh {
		a.CalculationMeanSTD(frequency, count)
		sigChan <- true
		if count.CountOfRecords == a.qualityOfValues {
			break
		}
	}

	fmt.Printf("Anomaly Detection. Mean: %.4f STD: %.4f k*STD: %.4f\n", a.Mean, a.STD, a.anomalyCoefficient*a.STD)

	for frequency := range frequencyCh {
		a.CalculationMeanSTD(frequency, count)
		sigChan <- true
	}
}

func NewAnomaliesDetector(anomalyCoefficient float64, qualityOfValues uint64) (*DetectorOfAnomalies, *Counts) {
	a := &DetectorOfAnomalies{
		Mean:               0,
		STD:                0,
		SumOfMean:          0,
		SumOfSTD:           0,
		anomalyCoefficient: anomalyCoefficient,
		Values:             make([]float64, 0, qualityOfValues),
		qualityOfValues:    qualityOfValues,
	}
	b := &Counts{
		CountOfRecords:   0,
		CountOfAnomalies: 0,
	}
	return a, b
}

func main() {
	anomalyCoefficient := flag.Float64("k", 0, "STD anomaly coefficient")
	flag.Parse()

	if flag.NArg() != 0 || *anomalyCoefficient == 0 {
		flag.Usage()
		os.Exit(1)
	}

	conn, err := grpc.Dial(consts.Port, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c := pb.NewTransmitterServiceClient(conn)

	stream, err := c.StreamData(context.Background(), &pb.StreamRequest{})
	if err != nil {
		log.Fatalf("could not stream data: %v", err)
	}

	db, err := models.InitDB()
	if err != nil {
		log.Fatalln(err)
	}

	frequencyCh := make(chan float64)
	sigChan := make(chan bool)
	defer close(frequencyCh)

	detectorOfAnomalies, counts := NewAnomaliesDetector(*anomalyCoefficient, qualityOfValues)
	detectorOfAnomalies.Db = db
	detectorOfAnomalies.Initialize(counts)

	go detectorOfAnomalies.DetectAnomalies(frequencyCh, sigChan, counts)

	for {
		var resp *pb.TransmitterData
		resp, err = stream.Recv()
		if err != nil {
			log.Fatalf("Error while receiving: %v", err)
		}

		log.Printf("Received message: session_id: %s, frequency: %f, timestamp: %d",
			resp.SessionId, resp.Frequency, resp.Timestamp)
		detectorOfAnomalies.Anomaly = models.Anomaly{SessionId: resp.SessionId, Frequency: resp.Frequency, Timestamp: time.Unix(resp.Timestamp, 0)}
		frequencyCh <- resp.Frequency
		<-sigChan
	}
}
