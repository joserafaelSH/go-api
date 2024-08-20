package healthservice

type HealthService struct {
    HealthMessage string
}

func CreateHealthService() *HealthService {
    return &HealthService{HealthMessage: "Server is running"}
}

func (h HealthService) getHealth() string {
    return h.HealthMessage
}

func (h HealthService) Execute(input any) string{
    return h.getHealth()
}
