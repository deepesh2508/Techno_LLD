package auth

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/redis/go-redis/v9"
)

type OTPService struct {
	redis *redis.Client
}

func NewOTPService(r *redis.Client) *OTPService {
	return &OTPService{redis: r}
}

func (o *OTPService) Send(email string) error {
	ctx := context.Background()

	otp := generateOTP()
	key := fmt.Sprintf("otp:%s", email)

	// store OTP with expiry
	err := o.redis.Set(ctx, key, otp, 5*time.Minute).Err()
	if err != nil {
		return err
	}

	// 🔥 Stub: real world me email / SMS gateway
	fmt.Printf("OTP for %s is %s\n", email, otp)

	return nil
}

func (o *OTPService) Verify(email string, otp string) bool {
	ctx := context.Background()
	key := fmt.Sprintf("otp:%s", email)

	val, err := o.redis.Get(ctx, key).Result()
	if err != nil {
		return false
	}

	if val != otp {
		return false
	}

	// OTP used → delete
	o.redis.Del(ctx, key)
	return true
}

func generateOTP() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}
