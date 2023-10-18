// version : V1
package models

type ISO8583Message struct {
	HeaderMsg                                  string `validate:"string,length=12"`
	MTI_Version                                int    `validate:"number,length=1"`
	MTI_MessageClass                           int    `validate:"number,length=1"`
	MTI_MessageFunction                        int    `validate:"number,length=1"`
	MTI_MessageOrigin                          int    `validate:"number,length=1"`
	PrimaryBitmap                              string `validate:"string,length=16"`
	SecondaryBitmap                            string `validate:"string,length=16"`
	ProcessingCode_003                         string `validate:"string,length=6"`
	AmountTransaction_004                      string `validate:"string,length=12"`
	TransmissionDatetime_007                   string `validate:"string,length=10"`
	SystemsTraceAuditNumber_011                string `validate:"string,length=6"`
	TimeLocalTransaction_012                   string `validate:"string,length=6"`
	DateLocalTransaction_013                   string `validate:"string,length=4"`
	DateCapture_017                            string `validate:"string,length=4"`
	PointOfServiceEntryMode_022                string `validate:"string,length=3"`
	AcquiringInstitutionIdentificationCode_032 string `validate:"string,length=11"`
	Track2Data_035                             string `validate:"string,length=37"`
	RetrievalReferenceNumber_037               string `validate:"string,length=12"`
	CardAcceptorTerminalIdentification_041     string `validate:"string,length=16"`
	CardAcceptorIdentificationCode_042         string `validate:"string,length=15"`
	CardAcceptorNameLocation_043               string `validate:"string,length=40"`
	AdditionalDataPrivate_048                  string `validate:"string,length=18"`
	CurrencyCodeTransaction_049                string `validate:"string,length=3"`
	AdviceReasonCode_060                       string `validate:"string,length=16"`
	ReceivingInstitutionIdentificationCode_100 string `validate:"string,length=11"`
	AccountIdentification1_102                 string `validate:"string,length=19"`
	ReservedPrivate_126                        string `validate:"string,length=16"`
}
