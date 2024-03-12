package database

type Store struct {
	ID   uint   `gorm:"primaryKey"`
	Guid string `gorm:"size:100"`
	Typ string `gorm:"size:100"`
	Desc string `gorm:"size:500"`
	Value1 string
	Value2 string
	Value3 string
	Value4 string   
}


type Keys struct {
	ID   uint   `gorm:"primaryKey"`
	Guid string `gorm:"size:100"`
	Typ string `gorm:"size:100"`
	Desc string `gorm:"size:500"`
	IdentKey string `gorm:"size:500"`
	PubKey string
	PrivKey string
	Fernet string
}

type Fusion struct {
	ID   uint   `gorm:"primaryKey"`
	SGuid string `gorm:"size:100"`
	KGuid string `gorm:"size:100"`
	Typ string `gorm:"size:100"`
	
}
// dies ist ein Test

