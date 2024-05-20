package middleware

import (
	"github.com/gin-gonic/gin"
	"src/ex02/models"
	"sync"
	"time"
)

var visitors = make(map[string]*models.Visitor)
var mtx sync.Mutex

func RateLimit(c *gin.Context) {
	mtx.Lock()
	defer mtx.Unlock()

	v, exists := visitors[c.ClientIP()]
	if !exists {
		visitors[c.ClientIP()] = &models.Visitor{
			LastSeen: time.Now(),
			Count:    1,
		}
	} else {
		if time.Since(v.LastSeen) > 1*time.Second {
			v.LastSeen = time.Now()
			v.Count = 0
		}

		v.Count++
		if v.Count > 100 {
			c.AbortWithStatusJSON(429, gin.H{"message": "Too many requests. Please try again later."})
			return
		}
	}

	c.Next()
}

/*В этом коде mtx.Lock() и defer mtx.Unlock() используются для обеспечения безопасности при одновременном
доступе к общим ресурсам в многопоточной среде. В данном случае общим ресурсом является карта visitors.

Когда несколько горутин (потоков выполнения в Go) пытаются одновременно изменить данные в карте visitors,
это может привести к состоянию гонки, что может вызвать непредсказуемое поведение программы.

mtx.Lock() блокирует мьютекс, тем самым предотвращая другие горутины от доступа к защищенной им области
памяти, пока текущая горутина не вызовет mtx.Unlock(). defer гарантирует, что mtx.Unlock() будет вызван
в конце функции RateLimit, даже если функция завершится досрочно из-за ошибки или вызова return.

Таким образом, эти вызовы обеспечивают безопасность при одновременном доступе к карте visitors, предотвращая
возможные состояния гонки.*/

/*В представленном вами коде горутины непосредственно не используются. Однако, веб-серверы, такие как тот,
что использует Gin, обычно обрабатывают каждый входящий HTTP-запрос в отдельной горутине. Это означает,
что если ваше промежуточное ПО RateLimit используется в таком контексте, функция RateLimit может быть
вызвана одновременно из нескольких горутин. Именно поэтому здесь используется мьютекс для синхронизации
доступа к общему ресурсу (visitors).*/
