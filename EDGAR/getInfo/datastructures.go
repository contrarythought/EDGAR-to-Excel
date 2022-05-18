package getInfo

import "time"

type CompanyTicker struct {
	CIK    int    `json:"cik_str"`
	Ticker string `json:"ticker"`
	Title  string `json:"title"`
}

type CompanyCollection struct {
	Collection map[string]CompanyTicker
}

type TickerToCIK struct {
	Collection map[string]string
}

type Submissions struct {
	CIK                               string   `json:"cik"`
	EntityType                        string   `json:"entityType"`
	SIC                               string   `json:"sic"`
	SICDescription                    string   `json:"sicDescription"`
	InsiderTransactionForOwnerExists  int      `json:"insiderTransactionForOwnerExists"`
	InsiderTransactionForIssuerExists int      `json:"insiderTransactionForIssuerExists"`
	Name                              string   `json:"name"`
	Tickers                           []string `json:"tickers"`
	Exchanges                         []string `json:"exchanges"`
	EIN                               string   `json:"ein"`
	Description                       string   `json:"description,omitempty"`
	Website                           string   `json:"website,omitempty"`
	InvestorWebsite                   string   `json:"investorWebsite,omitempty"`
	Category                          string   `json:"category"`
	FiscalYearEnd                     string   `json:"fiscalYearEnd"`
	StateOfIncorporation              string   `json:"stateOfIncorporation"`
	StateOfIncorporationDescription   string   `json:"stateOfIncorporationDescription"`
	Addresses                         struct {
		Mailing struct {
			Street1                   string      `json:"street1"`
			Street2                   interface{} `json:"street2"`
			City                      string      `json:"city"`
			StateOrCountry            string      `json:"stateOrCountry"`
			ZipCode                   string      `json:"zipCode"`
			StateOrCountryDescription string      `json:"stateOrCountryDescription"`
		} `json:"mailing"`
		Business struct {
			Street1                   string      `json:"street1"`
			Street2                   interface{} `json:"street2"`
			City                      string      `json:"city"`
			StateOrCountry            string      `json:"stateOrCountry"`
			ZipCode                   string      `json:"zipCode"`
			StateOrCountryDescription string      `json:"stateOrCountryDescription"`
		} `json:"business"`
	} `json:"addresses"`
	Phone       string `json:"phone"`
	Flags       string `json:"flags,omitempty"`
	FormerNames []struct {
		Name string    `json:"name"`
		From time.Time `json:"from"`
		To   time.Time `json:"to"`
	}
	Filings struct {
		Recent struct {
			AccessionNumber    []string    `json:"accessionNumber"`
			FilingDate         []string    `json:"filingDate"`
			ReportDate         []string    `json:"reportDate"`
			AcceptanceDateTime []time.Time `json:"acceptanceDateTime"`
			Act                []string    `json:"act"`
			Form               []string    `json:"form"`
			FileNumber         []string    `json:"fileNumber"`
			FilmNumber         []string    `json:"filmNumber"`
			Items              []string    `json:"items"`
			Size               []int       `json:"size"`
			IsXBRL             []int       `json:"isXBRL"`
			IsInlineXBRL       []int       `json:"isInlineXBRL"`
			PrimaryDocument    []string    `json:"primaryDocument"`
		} `json:"recent"`
		Files []struct {
			Name        string `json:"name"`
			FilingCount int    `json:"filingCount"`
			FilingFrom  string `json:"filingFrom"`
			FilingTo    string `json:"filingTo"`
		} `json:"files"`
	} `json:"filings"`
}

type CompanyFacts struct {
	Cik        int    `json:"cik"`
	EntityName string `json:"entityName"`
	Facts      struct {
		Dei struct {
			EntityCommonStockSharesOutstanding struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Shares []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"shares"`
				} `json:"units"`
			} `json:"EntityCommonStockSharesOutstanding"`
			EntityPublicFloat struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"EntityPublicFloat"`
		} `json:"dei"`
		UsGaap struct {
			AccountsPayable struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"AccountsPayable"`
			AccountsPayableCurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"AccountsPayableCurrent"`
			AccountsReceivableNetCurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"AccountsReceivableNetCurrent"`
			AccruedIncomeTaxesCurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"AccruedIncomeTaxesCurrent"`
			AccruedIncomeTaxesNoncurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"AccruedIncomeTaxesNoncurrent"`
			AccruedLiabilities struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"AccruedLiabilities"`
			AccruedLiabilitiesCurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"AccruedLiabilitiesCurrent"`
			AccruedMarketingCostsCurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"AccruedMarketingCostsCurrent"`
			AccumulatedDepreciationDepletionAndAmortizationPropertyPlantAndEquipment struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"AccumulatedDepreciationDepletionAndAmortizationPropertyPlantAndEquipment"`
			AccumulatedOtherComprehensiveIncomeLossAvailableForSaleSecuritiesAdjustmentNetOfTax struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"AccumulatedOtherComprehensiveIncomeLossAvailableForSaleSecuritiesAdjustmentNetOfTax"`
			AccumulatedOtherComprehensiveIncomeLossCumulativeChangesInNetGainLossFromCashFlowHedgesEffectNetOfTax struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"AccumulatedOtherComprehensiveIncomeLossCumulativeChangesInNetGainLossFromCashFlowHedgesEffectNetOfTax"`
			AccumulatedOtherComprehensiveIncomeLossForeignCurrencyTranslationAdjustmentNetOfTax struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"AccumulatedOtherComprehensiveIncomeLossForeignCurrencyTranslationAdjustmentNetOfTax"`
			AccumulatedOtherComprehensiveIncomeLossNetOfTax struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"AccumulatedOtherComprehensiveIncomeLossNetOfTax"`
			AdjustmentsToAdditionalPaidInCapitalSharebasedCompensationRequisiteServicePeriodRecognitionValue struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"AdjustmentsToAdditionalPaidInCapitalSharebasedCompensationRequisiteServicePeriodRecognitionValue"`
			AdjustmentsToAdditionalPaidInCapitalTaxEffectFromShareBasedCompensation struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"AdjustmentsToAdditionalPaidInCapitalTaxEffectFromShareBasedCompensation"`
			AdvertisingExpense struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"AdvertisingExpense"`
			AllocatedShareBasedCompensationExpense struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"AllocatedShareBasedCompensationExpense"`
			AllowanceForDoubtfulAccountsReceivableCurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"AllowanceForDoubtfulAccountsReceivableCurrent"`
			AmortizationOfIntangibleAssets struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"AmortizationOfIntangibleAssets"`
			AntidilutiveSecuritiesExcludedFromComputationOfEarningsPerShareAmount struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Shares []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"shares"`
				} `json:"units"`
			} `json:"AntidilutiveSecuritiesExcludedFromComputationOfEarningsPerShareAmount"`
			Assets struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"Assets"`
			AssetsCurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"AssetsCurrent"`
			AssetsNoncurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"AssetsNoncurrent"`
			AvailableForSaleDebtSecuritiesAccumulatedGrossUnrealizedGainBeforeTax struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"AvailableForSaleDebtSecuritiesAccumulatedGrossUnrealizedGainBeforeTax"`
			AvailableForSaleDebtSecuritiesAccumulatedGrossUnrealizedLossBeforeTax struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"AvailableForSaleDebtSecuritiesAccumulatedGrossUnrealizedLossBeforeTax"`
			AvailableForSaleDebtSecuritiesAmortizedCostBasis struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"AvailableForSaleDebtSecuritiesAmortizedCostBasis"`
			AvailableForSaleSecurities struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"AvailableForSaleSecurities"`
			AvailableForSaleSecuritiesAccumulatedGrossUnrealizedGainBeforeTax struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"AvailableForSaleSecuritiesAccumulatedGrossUnrealizedGainBeforeTax"`
			AvailableForSaleSecuritiesAccumulatedGrossUnrealizedLossBeforeTax struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"AvailableForSaleSecuritiesAccumulatedGrossUnrealizedLossBeforeTax"`
			AvailableForSaleSecuritiesAmortizedCost struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"AvailableForSaleSecuritiesAmortizedCost"`
			AvailableForSaleSecuritiesContinuousUnrealizedLossPosition12MonthsOrLongerAccumulatedLoss struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"AvailableForSaleSecuritiesContinuousUnrealizedLossPosition12MonthsOrLongerAccumulatedLoss"`
			AvailableForSaleSecuritiesContinuousUnrealizedLossPositionAccumulatedLoss struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"AvailableForSaleSecuritiesContinuousUnrealizedLossPositionAccumulatedLoss"`
			AvailableForSaleSecuritiesContinuousUnrealizedLossPositionFairValue struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"AvailableForSaleSecuritiesContinuousUnrealizedLossPositionFairValue"`
			AvailableForSaleSecuritiesContinuousUnrealizedLossPositionLessThan12MonthsAccumulatedLoss struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"AvailableForSaleSecuritiesContinuousUnrealizedLossPositionLessThan12MonthsAccumulatedLoss"`
			AvailableForSaleSecuritiesContinuousUnrealizedLossPositionLessThanTwelveMonthsFairValue struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"AvailableForSaleSecuritiesContinuousUnrealizedLossPositionLessThanTwelveMonthsFairValue"`
			AvailableForSaleSecuritiesContinuousUnrealizedLossPositionTwelveMonthsOrLongerFairValue struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"AvailableForSaleSecuritiesContinuousUnrealizedLossPositionTwelveMonthsOrLongerFairValue"`
			AvailableForSaleSecuritiesCurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"AvailableForSaleSecuritiesCurrent"`
			AvailableForSaleSecuritiesDebtMaturitiesRollingAfterYearTenFairValue struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"AvailableForSaleSecuritiesDebtMaturitiesRollingAfterYearTenFairValue"`
			AvailableForSaleSecuritiesDebtMaturitiesRollingYearSixThroughTenFairValue struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"AvailableForSaleSecuritiesDebtMaturitiesRollingYearSixThroughTenFairValue"`
			AvailableForSaleSecuritiesDebtMaturitiesRollingYearTwoThroughFiveFairValue struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"AvailableForSaleSecuritiesDebtMaturitiesRollingYearTwoThroughFiveFairValue"`
			AvailableForSaleSecuritiesDebtMaturitiesSingleMaturityDate struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"AvailableForSaleSecuritiesDebtMaturitiesSingleMaturityDate"`
			AvailableForSaleSecuritiesDebtSecurities struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"AvailableForSaleSecuritiesDebtSecurities"`
			AvailableForSaleSecuritiesDebtSecuritiesCurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"AvailableForSaleSecuritiesDebtSecuritiesCurrent"`
			AvailableForSaleSecuritiesDebtSecuritiesNoncurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"AvailableForSaleSecuritiesDebtSecuritiesNoncurrent"`
			AvailableForSaleSecuritiesFairValueDisclosure struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"AvailableForSaleSecuritiesFairValueDisclosure"`
			AvailableForSaleSecuritiesGrossUnrealizedLosses1 struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"AvailableForSaleSecuritiesGrossUnrealizedLosses1"`
			AvailableForSaleSecuritiesNoncurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"AvailableForSaleSecuritiesNoncurrent"`
			AvailableforsaleSecuritiesGrossUnrealizedGain struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"AvailableforsaleSecuritiesGrossUnrealizedGain"`
			BusinessAcquisitionCostOfAcquiredEntityPurchasePrice struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"BusinessAcquisitionCostOfAcquiredEntityPurchasePrice"`
			BusinessAcquisitionPurchasePriceAllocationAmortizableIntangibleAssets struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"BusinessAcquisitionPurchasePriceAllocationAmortizableIntangibleAssets"`
			BusinessAcquisitionPurchasePriceAllocationGoodwillAmount struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"BusinessAcquisitionPurchasePriceAllocationGoodwillAmount"`
			BusinessAcquisitionPurchasePriceAllocationLiabilitiesAssumed struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"BusinessAcquisitionPurchasePriceAllocationLiabilitiesAssumed"`
			BusinessCombinationRecognizedIdentifiableAssetsAcquiredAndLiabilitiesAssumedIntangibleAssetsOtherThanGoodwill struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"BusinessCombinationRecognizedIdentifiableAssetsAcquiredAndLiabilitiesAssumedIntangibleAssetsOtherThanGoodwill"`
			BusinessCombinationRecognizedIdentifiableAssetsAcquiredAndLiabilitiesAssumedLiabilities struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"BusinessCombinationRecognizedIdentifiableAssetsAcquiredAndLiabilitiesAssumedLiabilities"`
			CapitalExpendituresIncurredButNotYetPaid struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"CapitalExpendituresIncurredButNotYetPaid"`
			CapitalizedComputerSoftwareAdditions struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"CapitalizedComputerSoftwareAdditions"`
			CapitalizedComputerSoftwareAmortization struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"CapitalizedComputerSoftwareAmortization"`
			Cash struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"Cash"`
			CashAndCashEquivalentsAtCarryingValue struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"CashAndCashEquivalentsAtCarryingValue"`
			CashAndCashEquivalentsPeriodIncreaseDecrease struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"CashAndCashEquivalentsPeriodIncreaseDecrease"`
			CashCashEquivalentsRestrictedCashAndRestrictedCashEquivalents struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"CashCashEquivalentsRestrictedCashAndRestrictedCashEquivalents"`
			CashCashEquivalentsRestrictedCashAndRestrictedCashEquivalentsPeriodIncreaseDecreaseIncludingExchangeRateEffect struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"CashCashEquivalentsRestrictedCashAndRestrictedCashEquivalentsPeriodIncreaseDecreaseIncludingExchangeRateEffect"`
			CashEquivalentsAtCarryingValue struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"CashEquivalentsAtCarryingValue"`
			ChangeInUnrealizedGainLossOnFairValueHedgingInstruments struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ChangeInUnrealizedGainLossOnFairValueHedgingInstruments"`
			CollateralAlreadyPostedAggregateFairValue struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"CollateralAlreadyPostedAggregateFairValue"`
			CommercialPaper struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"CommercialPaper"`
			CommitmentsAndContingencies struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"CommitmentsAndContingencies"`
			CommonStockDividendsPerShareCashPaid struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					USDShares []struct {
						Start string  `json:"start"`
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
					} `json:"USD/shares"`
				} `json:"units"`
			} `json:"CommonStockDividendsPerShareCashPaid"`
			CommonStockDividendsPerShareDeclared struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					USDShares []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD/shares"`
				} `json:"units"`
			} `json:"CommonStockDividendsPerShareDeclared"`
			CommonStockNoParValue struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					USDShares []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD/shares"`
				} `json:"units"`
			} `json:"CommonStockNoParValue"`
			CommonStockParOrStatedValuePerShare struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					USDShares []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD/shares"`
				} `json:"units"`
			} `json:"CommonStockParOrStatedValuePerShare"`
			CommonStockSharesAuthorized struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Shares []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"shares"`
				} `json:"units"`
			} `json:"CommonStockSharesAuthorized"`
			CommonStockSharesIssued struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Shares []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"shares"`
				} `json:"units"`
			} `json:"CommonStockSharesIssued"`
			CommonStockSharesOutstanding struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Shares []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"shares"`
				} `json:"units"`
			} `json:"CommonStockSharesOutstanding"`
			CommonStocksIncludingAdditionalPaidInCapital struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"CommonStocksIncludingAdditionalPaidInCapital"`
			CommonStockValue struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"CommonStockValue"`
			ComprehensiveIncomeNetOfTax struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ComprehensiveIncomeNetOfTax"`
			ContractWithCustomerLiability struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ContractWithCustomerLiability"`
			ContractWithCustomerLiabilityCurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ContractWithCustomerLiabilityCurrent"`
			ContractWithCustomerLiabilityRevenueRecognized struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ContractWithCustomerLiabilityRevenueRecognized"`
			CostOfGoodsAndServicesSold struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"CostOfGoodsAndServicesSold"`
			CostOfGoodsAndServicesSoldDepreciation struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"CostOfGoodsAndServicesSoldDepreciation"`
			CostOfReimbursableExpense struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"CostOfReimbursableExpense"`
			CumulativeEffectOfInitialAdoptionOfFIN48 struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"CumulativeEffectOfInitialAdoptionOfFIN48"`
			CumulativeEffectOfInitialAdoptionOfNewAccountingPrinciple struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"CumulativeEffectOfInitialAdoptionOfNewAccountingPrinciple"`
			CurrentFederalTaxExpenseBenefit struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"CurrentFederalTaxExpenseBenefit"`
			CurrentForeignTaxExpenseBenefit struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"CurrentForeignTaxExpenseBenefit"`
			CurrentStateAndLocalTaxExpenseBenefit struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"CurrentStateAndLocalTaxExpenseBenefit"`
			DebtInstrumentCarryingAmount struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DebtInstrumentCarryingAmount"`
			DebtInstrumentUnamortizedDiscount struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DebtInstrumentUnamortizedDiscount"`
			DebtInstrumentUnamortizedDiscountPremiumAndDebtIssuanceCostsNet struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DebtInstrumentUnamortizedDiscountPremiumAndDebtIssuanceCostsNet"`
			DebtInstrumentUnamortizedDiscountPremiumNet struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DebtInstrumentUnamortizedDiscountPremiumNet"`
			DebtSecuritiesAvailableForSaleContinuousUnrealizedLossPosition12MonthsOrLonger struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DebtSecuritiesAvailableForSaleContinuousUnrealizedLossPosition12MonthsOrLonger"`
			DebtSecuritiesAvailableForSaleContinuousUnrealizedLossPosition12MonthsOrLongerAccumulatedLoss struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DebtSecuritiesAvailableForSaleContinuousUnrealizedLossPosition12MonthsOrLongerAccumulatedLoss"`
			DebtSecuritiesAvailableForSaleContinuousUnrealizedLossPositionLessThan12Months struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DebtSecuritiesAvailableForSaleContinuousUnrealizedLossPositionLessThan12Months"`
			DebtSecuritiesAvailableForSaleContinuousUnrealizedLossPositionLessThan12MonthsAccumulatedLoss struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DebtSecuritiesAvailableForSaleContinuousUnrealizedLossPositionLessThan12MonthsAccumulatedLoss"`
			DebtSecuritiesAvailableForSaleUnrealizedLossPosition struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DebtSecuritiesAvailableForSaleUnrealizedLossPosition"`
			DebtSecuritiesAvailableForSaleUnrealizedLossPositionAccumulatedLoss struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DebtSecuritiesAvailableForSaleUnrealizedLossPositionAccumulatedLoss"`
			DecreaseInUnrecognizedTaxBenefitsIsReasonablyPossible struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DecreaseInUnrecognizedTaxBenefitsIsReasonablyPossible"`
			DeferredFederalIncomeTaxExpenseBenefit struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DeferredFederalIncomeTaxExpenseBenefit"`
			DeferredForeignIncomeTaxExpenseBenefit struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DeferredForeignIncomeTaxExpenseBenefit"`
			DeferredIncomeTaxExpenseBenefit struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DeferredIncomeTaxExpenseBenefit"`
			DeferredIncomeTaxLiabilities struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DeferredIncomeTaxLiabilities"`
			DeferredIncomeTaxLiabilitiesNet struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DeferredIncomeTaxLiabilitiesNet"`
			DeferredRevenueCurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DeferredRevenueCurrent"`
			DeferredRevenueNoncurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DeferredRevenueNoncurrent"`
			DeferredStateAndLocalIncomeTaxExpenseBenefit struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DeferredStateAndLocalIncomeTaxExpenseBenefit"`
			DeferredTaxAssetsDeferredIncome struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DeferredTaxAssetsDeferredIncome"`
			DeferredTaxAssetsGoodwillAndIntangibleAssets struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DeferredTaxAssetsGoodwillAndIntangibleAssets"`
			DeferredTaxAssetsGross struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DeferredTaxAssetsGross"`
			DeferredTaxAssetsLiabilitiesNet struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DeferredTaxAssetsLiabilitiesNet"`
			DeferredTaxAssetsLiabilitiesNetCurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DeferredTaxAssetsLiabilitiesNetCurrent"`
			DeferredTaxAssetsNet struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DeferredTaxAssetsNet"`
			DeferredTaxAssetsNetCurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DeferredTaxAssetsNetCurrent"`
			DeferredTaxAssetsOther struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DeferredTaxAssetsOther"`
			DeferredTaxAssetsPropertyPlantAndEquipment struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DeferredTaxAssetsPropertyPlantAndEquipment"`
			DeferredTaxAssetsTaxCreditCarryforwards struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DeferredTaxAssetsTaxCreditCarryforwards"`
			DeferredTaxAssetsTaxCreditCarryforwardsForeign struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DeferredTaxAssetsTaxCreditCarryforwardsForeign"`
			DeferredTaxAssetsTaxCreditCarryforwardsResearch struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DeferredTaxAssetsTaxCreditCarryforwardsResearch"`
			DeferredTaxAssetsTaxDeferredExpenseCompensationAndBenefitsShareBasedCompensationCost struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DeferredTaxAssetsTaxDeferredExpenseCompensationAndBenefitsShareBasedCompensationCost"`
			DeferredTaxAssetsTaxDeferredExpenseReservesAndAccruals struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DeferredTaxAssetsTaxDeferredExpenseReservesAndAccruals"`
			DeferredTaxAssetsUnrealizedLossesOnAvailableforSaleSecuritiesGross struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DeferredTaxAssetsUnrealizedLossesOnAvailableforSaleSecuritiesGross"`
			DeferredTaxAssetsValuationAllowance struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DeferredTaxAssetsValuationAllowance"`
			DeferredTaxLiabilities struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DeferredTaxLiabilities"`
			DeferredTaxLiabilitiesLeasingArrangements struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DeferredTaxLiabilitiesLeasingArrangements"`
			DeferredTaxLiabilitiesNoncurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DeferredTaxLiabilitiesNoncurrent"`
			DeferredTaxLiabilitiesOther struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DeferredTaxLiabilitiesOther"`
			DeferredTaxLiabilitiesOtherComprehensiveIncome struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DeferredTaxLiabilitiesOtherComprehensiveIncome"`
			DeferredTaxLiabilitiesUndistributedForeignEarnings struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DeferredTaxLiabilitiesUndistributedForeignEarnings"`
			DeferredTaxLiabilityNotRecognizedAmountOfUnrecognizedDeferredTaxLiabilityUndistributedEarningsOfForeignSubsidiaries struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DeferredTaxLiabilityNotRecognizedAmountOfUnrecognizedDeferredTaxLiabilityUndistributedEarningsOfForeignSubsidiaries"`
			DefinedContributionPlanCostRecognized struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DefinedContributionPlanCostRecognized"`
			DefinedContributionPlanMaximumAnnualContributionPerEmployeeAmount struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DefinedContributionPlanMaximumAnnualContributionPerEmployeeAmount"`
			DefinedContributionPlanMaximumAnnualContributionsPerEmployeeAmount struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DefinedContributionPlanMaximumAnnualContributionsPerEmployeeAmount"`
			Depreciation struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"Depreciation"`
			DepreciationAmortizationAndAccretionNet struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DepreciationAmortizationAndAccretionNet"`
			DepreciationAndAmortization struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DepreciationAndAmortization"`
			DepreciationDepletionAndAmortization struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DepreciationDepletionAndAmortization"`
			DerivativeCollateralObligationToReturnCash struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DerivativeCollateralObligationToReturnCash"`
			DerivativeFairValueOfDerivativeNet struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DerivativeFairValueOfDerivativeNet"`
			DerivativeInstrumentsGainLossReclassifiedFromAccumulatedOCIIntoIncomeEffectivePortionNet struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DerivativeInstrumentsGainLossReclassifiedFromAccumulatedOCIIntoIncomeEffectivePortionNet"`
			DerivativeInstrumentsGainLossRecognizedInIncomeIneffectivePortionAndAmountExcludedFromEffectivenessTestingNet struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DerivativeInstrumentsGainLossRecognizedInIncomeIneffectivePortionAndAmountExcludedFromEffectivenessTestingNet"`
			DerivativeInstrumentsGainLossRecognizedInOtherComprehensiveIncomeEffectivePortionNet struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DerivativeInstrumentsGainLossRecognizedInOtherComprehensiveIncomeEffectivePortionNet"`
			DerivativeInstrumentsNotDesignatedAsHedgingInstrumentsGainLossNet struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"DerivativeInstrumentsNotDesignatedAsHedgingInstrumentsGainLossNet"`
			Dividends struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"Dividends"`
			EarningsPerShareBasic struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					USDShares []struct {
						Start string  `json:"start"`
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame,omitempty"`
					} `json:"USD/shares"`
				} `json:"units"`
			} `json:"EarningsPerShareBasic"`
			EarningsPerShareDiluted struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					USDShares []struct {
						Start string  `json:"start"`
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame,omitempty"`
					} `json:"USD/shares"`
				} `json:"units"`
			} `json:"EarningsPerShareDiluted"`
			EffectiveIncomeTaxRateContinuingOperations struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Pure []struct {
						Start string  `json:"start"`
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame,omitempty"`
					} `json:"pure"`
				} `json:"units"`
			} `json:"EffectiveIncomeTaxRateContinuingOperations"`
			EffectiveIncomeTaxRateReconciliationAtFederalStatutoryIncomeTaxRate struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Pure []struct {
						Start string  `json:"start,omitempty"`
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame,omitempty"`
					} `json:"pure"`
				} `json:"units"`
			} `json:"EffectiveIncomeTaxRateReconciliationAtFederalStatutoryIncomeTaxRate"`
			EffectiveIncomeTaxRateReconciliationForeignIncomeTaxRateDifferential struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Pure []struct {
						Start string  `json:"start"`
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame,omitempty"`
					} `json:"pure"`
				} `json:"units"`
			} `json:"EffectiveIncomeTaxRateReconciliationForeignIncomeTaxRateDifferential"`
			EffectiveIncomeTaxRateReconciliationShareBasedCompensationExcessTaxBenefitAmount struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"EffectiveIncomeTaxRateReconciliationShareBasedCompensationExcessTaxBenefitAmount"`
			EffectiveIncomeTaxRateReconciliationTaxCutsAndJobsActOf2017Amount struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"EffectiveIncomeTaxRateReconciliationTaxCutsAndJobsActOf2017Amount"`
			EmployeeRelatedLiabilitiesCurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"EmployeeRelatedLiabilitiesCurrent"`
			EmployeeServiceShareBasedCompensationCashFlowEffectCashUsedToSettleAwards struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"EmployeeServiceShareBasedCompensationCashFlowEffectCashUsedToSettleAwards"`
			EmployeeServiceShareBasedCompensationNonvestedAwardsTotalCompensationCostNotYetRecognized struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"EmployeeServiceShareBasedCompensationNonvestedAwardsTotalCompensationCostNotYetRecognized"`
			EmployeeServiceShareBasedCompensationNonvestedAwardsTotalCompensationCostNotYetRecognizedPeriodForRecognition struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Year []struct {
						Start string  `json:"start"`
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame,omitempty"`
					} `json:"Year"`
				} `json:"units"`
			} `json:"EmployeeServiceShareBasedCompensationNonvestedAwardsTotalCompensationCostNotYetRecognizedPeriodForRecognition"`
			EmployeeServiceShareBasedCompensationTaxBenefitFromCompensationExpense struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"EmployeeServiceShareBasedCompensationTaxBenefitFromCompensationExpense"`
			EmployeeServiceShareBasedCompensationUnrecognizedCompensationCostsOnNonvestedAwards struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"EmployeeServiceShareBasedCompensationUnrecognizedCompensationCostsOnNonvestedAwards"`
			EmployeeServiceShareBasedCompensationUnrecognizedCompensationCostsOnNonvestedAwardsWeightedAveragePeriodOfRecognition struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Pure []struct {
						Start string  `json:"start"`
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame"`
					} `json:"pure"`
				} `json:"units"`
			} `json:"EmployeeServiceShareBasedCompensationUnrecognizedCompensationCostsOnNonvestedAwardsWeightedAveragePeriodOfRecognition"`
			EquitySecuritiesWithoutReadilyDeterminableFairValueAmount struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"EquitySecuritiesWithoutReadilyDeterminableFairValueAmount"`
			ExcessTaxBenefitFromShareBasedCompensationFinancingActivities struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ExcessTaxBenefitFromShareBasedCompensationFinancingActivities"`
			FairValueAssetsMeasuredOnRecurringBasisCashAndCashEquivalents struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"FairValueAssetsMeasuredOnRecurringBasisCashAndCashEquivalents"`
			FairValueAssetsMeasuredOnRecurringBasisDerivativeFinancialInstrumentsAssets struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"FairValueAssetsMeasuredOnRecurringBasisDerivativeFinancialInstrumentsAssets"`
			FairValueHedgesAtFairValueNet struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"FairValueHedgesAtFairValueNet"`
			FairValueLiabilitiesMeasuredOnRecurringBasisDerivativeFinancialInstrumentsLiabilities struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"FairValueLiabilitiesMeasuredOnRecurringBasisDerivativeFinancialInstrumentsLiabilities"`
			FederalIncomeTaxExpenseBenefitContinuingOperations struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"FederalIncomeTaxExpenseBenefitContinuingOperations"`
			FinanceLeaseLiability struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"FinanceLeaseLiability"`
			FinanceLeaseLiabilityCurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"FinanceLeaseLiabilityCurrent"`
			FinanceLeaseLiabilityNoncurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"FinanceLeaseLiabilityNoncurrent"`
			FinanceLeaseLiabilityPaymentsDue struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"FinanceLeaseLiabilityPaymentsDue"`
			FinanceLeaseLiabilityPaymentsDueAfterYearFive struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"FinanceLeaseLiabilityPaymentsDueAfterYearFive"`
			FinanceLeaseLiabilityPaymentsDueNextTwelveMonths struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"FinanceLeaseLiabilityPaymentsDueNextTwelveMonths"`
			FinanceLeaseLiabilityPaymentsDueYearFive struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"FinanceLeaseLiabilityPaymentsDueYearFive"`
			FinanceLeaseLiabilityPaymentsDueYearFour struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"FinanceLeaseLiabilityPaymentsDueYearFour"`
			FinanceLeaseLiabilityPaymentsDueYearThree struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"FinanceLeaseLiabilityPaymentsDueYearThree"`
			FinanceLeaseLiabilityPaymentsDueYearTwo struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"FinanceLeaseLiabilityPaymentsDueYearTwo"`
			FinanceLeaseLiabilityPaymentsRemainderOfFiscalYear struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"FinanceLeaseLiabilityPaymentsRemainderOfFiscalYear"`
			FinanceLeaseLiabilityUndiscountedExcessAmount struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"FinanceLeaseLiabilityUndiscountedExcessAmount"`
			FinanceLeaseRightOfUseAsset struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"FinanceLeaseRightOfUseAsset"`
			FiniteLivedIntangibleAssetsAccumulatedAmortization struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"FiniteLivedIntangibleAssetsAccumulatedAmortization"`
			FiniteLivedIntangibleAssetsAmortizationExpense struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"FiniteLivedIntangibleAssetsAmortizationExpense"`
			FiniteLivedIntangibleAssetsAmortizationExpenseAfterYearFive struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"FiniteLivedIntangibleAssetsAmortizationExpenseAfterYearFive"`
			FiniteLivedIntangibleAssetsAmortizationExpenseNextTwelveMonths struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"FiniteLivedIntangibleAssetsAmortizationExpenseNextTwelveMonths"`
			FiniteLivedIntangibleAssetsAmortizationExpenseYearFive struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"FiniteLivedIntangibleAssetsAmortizationExpenseYearFive"`
			FiniteLivedIntangibleAssetsAmortizationExpenseYearFour struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"FiniteLivedIntangibleAssetsAmortizationExpenseYearFour"`
			FiniteLivedIntangibleAssetsAmortizationExpenseYearThree struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"FiniteLivedIntangibleAssetsAmortizationExpenseYearThree"`
			FiniteLivedIntangibleAssetsAmortizationExpenseYearTwo struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"FiniteLivedIntangibleAssetsAmortizationExpenseYearTwo"`
			FiniteLivedIntangibleAssetsGross struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"FiniteLivedIntangibleAssetsGross"`
			FiniteLivedIntangibleAssetsNet struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"FiniteLivedIntangibleAssetsNet"`
			FiniteLivedIntangibleAssetsUsefulLifeMaximum struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Pure []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"pure"`
					Year []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"Year"`
				} `json:"units"`
			} `json:"FiniteLivedIntangibleAssetsUsefulLifeMaximum"`
			FiniteLivedIntangibleAssetsUsefulLifeMinimum struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Pure []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"pure"`
					Year []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"Year"`
				} `json:"units"`
			} `json:"FiniteLivedIntangibleAssetsUsefulLifeMinimum"`
			ForeignCurrencyDerivativeAssetsAtFairValue struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ForeignCurrencyDerivativeAssetsAtFairValue"`
			ForeignCurrencyDerivativeInstrumentsNotDesignatedAsHedgingInstrumentsAssetAtFairValue struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ForeignCurrencyDerivativeInstrumentsNotDesignatedAsHedgingInstrumentsAssetAtFairValue"`
			ForeignCurrencyDerivativeInstrumentsNotDesignatedAsHedgingInstrumentsLiabilityAtFairValue struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ForeignCurrencyDerivativeInstrumentsNotDesignatedAsHedgingInstrumentsLiabilityAtFairValue"`
			ForeignCurrencyDerivativeLiabilitiesAtFairValue struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ForeignCurrencyDerivativeLiabilitiesAtFairValue"`
			ForeignIncomeTaxExpenseBenefitContinuingOperations struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ForeignIncomeTaxExpenseBenefitContinuingOperations"`
			FurnitureAndFixturesGross struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"FurnitureAndFixturesGross"`
			FutureAmortizationExpenseYearFive struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"FutureAmortizationExpenseYearFive"`
			FutureAmortizationExpenseYearFour struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"FutureAmortizationExpenseYearFour"`
			FutureAmortizationExpenseYearOne struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"FutureAmortizationExpenseYearOne"`
			FutureAmortizationExpenseYearThree struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"FutureAmortizationExpenseYearThree"`
			FutureAmortizationExpenseYearTwo struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"FutureAmortizationExpenseYearTwo"`
			GainLossFromComponentsExcludedFromAssessmentOfFairValueHedgeEffectivenessNet struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"GainLossFromComponentsExcludedFromAssessmentOfFairValueHedgeEffectivenessNet"`
			GainLossOnSaleOfPropertyPlantEquipment struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"GainLossOnSaleOfPropertyPlantEquipment"`
			Goodwill struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"Goodwill"`
			GoodwillAcquiredDuringPeriod struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"GoodwillAcquiredDuringPeriod"`
			GoodwillImpairmentLoss struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"GoodwillImpairmentLoss"`
			GrossProfit struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"GrossProfit"`
			HeldToMaturitySecurities struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"HeldToMaturitySecurities"`
			ImpairmentOfIntangibleAssetsExcludingGoodwill struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ImpairmentOfIntangibleAssetsExcludingGoodwill"`
			ImpairmentOfIntangibleAssetsIndefinitelivedExcludingGoodwill struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ImpairmentOfIntangibleAssetsIndefinitelivedExcludingGoodwill"`
			IncomeLossFromContinuingOperationsBeforeIncomeTaxesExtraordinaryItemsNoncontrollingInterest struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"IncomeLossFromContinuingOperationsBeforeIncomeTaxesExtraordinaryItemsNoncontrollingInterest"`
			IncomeLossFromContinuingOperationsBeforeIncomeTaxesForeign struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"IncomeLossFromContinuingOperationsBeforeIncomeTaxesForeign"`
			IncomeLossFromContinuingOperationsBeforeIncomeTaxesMinorityInterestAndIncomeLossFromEquityMethodInvestments struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"IncomeLossFromContinuingOperationsBeforeIncomeTaxesMinorityInterestAndIncomeLossFromEquityMethodInvestments"`
			IncomeTaxExaminationInterestExpense struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"IncomeTaxExaminationInterestExpense"`
			IncomeTaxExpenseBenefit struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"IncomeTaxExpenseBenefit"`
			IncomeTaxesPaidNet struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"IncomeTaxesPaidNet"`
			IncomeTaxReconciliationDeductionsQualifiedProductionActivities struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"IncomeTaxReconciliationDeductionsQualifiedProductionActivities"`
			IncomeTaxReconciliationForeignIncomeTaxRateDifferential struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"IncomeTaxReconciliationForeignIncomeTaxRateDifferential"`
			IncomeTaxReconciliationIncomeTaxExpenseBenefitAtFederalStatutoryIncomeTaxRate struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"IncomeTaxReconciliationIncomeTaxExpenseBenefitAtFederalStatutoryIncomeTaxRate"`
			IncomeTaxReconciliationNondeductibleExpenseShareBasedCompensationCost struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"IncomeTaxReconciliationNondeductibleExpenseShareBasedCompensationCost"`
			IncomeTaxReconciliationOtherAdjustments struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"IncomeTaxReconciliationOtherAdjustments"`
			IncomeTaxReconciliationOtherReconcilingItems struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"IncomeTaxReconciliationOtherReconcilingItems"`
			IncomeTaxReconciliationStateAndLocalIncomeTaxes struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"IncomeTaxReconciliationStateAndLocalIncomeTaxes"`
			IncomeTaxReconciliationTaxContingenciesDomestic struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"IncomeTaxReconciliationTaxContingenciesDomestic"`
			IncomeTaxReconciliationTaxCreditsResearch struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"IncomeTaxReconciliationTaxCreditsResearch"`
			IncomeTaxReconciliationTaxSettlementsDomestic struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"IncomeTaxReconciliationTaxSettlementsDomestic"`
			IncreaseDecreaseInAccountsPayable struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"IncreaseDecreaseInAccountsPayable"`
			IncreaseDecreaseInAccountsReceivable struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"IncreaseDecreaseInAccountsReceivable"`
			IncreaseDecreaseInContractWithCustomerLiability struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"IncreaseDecreaseInContractWithCustomerLiability"`
			IncreaseDecreaseInDeferredRevenue struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"IncreaseDecreaseInDeferredRevenue"`
			IncreaseDecreaseInInventories struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"IncreaseDecreaseInInventories"`
			IncreaseDecreaseInOtherOperatingAssets struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"IncreaseDecreaseInOtherOperatingAssets"`
			IncreaseDecreaseInOtherOperatingLiabilities struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"IncreaseDecreaseInOtherOperatingLiabilities"`
			IncreaseDecreaseInOtherReceivables struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"IncreaseDecreaseInOtherReceivables"`
			IndefiniteLivedIntangibleAssetsExcludingGoodwill struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"IndefiniteLivedIntangibleAssetsExcludingGoodwill"`
			IntangibleAssetsGrossExcludingGoodwill struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"IntangibleAssetsGrossExcludingGoodwill"`
			IntangibleAssetsNetExcludingGoodwill struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"IntangibleAssetsNetExcludingGoodwill"`
			InterestCostsIncurred struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"InterestCostsIncurred"`
			InterestExpense struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"InterestExpense"`
			InterestExpenseDebt struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"InterestExpenseDebt"`
			InterestPaid struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"InterestPaid"`
			InterestPaidNet struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"InterestPaidNet"`
			InventoryFinishedGoodsNetOfReserves struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"InventoryFinishedGoodsNetOfReserves"`
			InventoryNet struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"InventoryNet"`
			InventoryPartsAndComponentsNetOfReserves struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"InventoryPartsAndComponentsNetOfReserves"`
			InvestmentIncomeInterestAndDividend struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"InvestmentIncomeInterestAndDividend"`
			LeaseAndRentalExpense struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"LeaseAndRentalExpense"`
			LeaseholdImprovementsGross struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"LeaseholdImprovementsGross"`
			LesseeOperatingLeaseLiabilityPaymentsDue struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"LesseeOperatingLeaseLiabilityPaymentsDue"`
			LesseeOperatingLeaseLiabilityPaymentsDueAfterYearFive struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"LesseeOperatingLeaseLiabilityPaymentsDueAfterYearFive"`
			LesseeOperatingLeaseLiabilityPaymentsDueNextTwelveMonths struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"LesseeOperatingLeaseLiabilityPaymentsDueNextTwelveMonths"`
			LesseeOperatingLeaseLiabilityPaymentsDueYearFive struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"LesseeOperatingLeaseLiabilityPaymentsDueYearFive"`
			LesseeOperatingLeaseLiabilityPaymentsDueYearFour struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"LesseeOperatingLeaseLiabilityPaymentsDueYearFour"`
			LesseeOperatingLeaseLiabilityPaymentsDueYearThree struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"LesseeOperatingLeaseLiabilityPaymentsDueYearThree"`
			LesseeOperatingLeaseLiabilityPaymentsDueYearTwo struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"LesseeOperatingLeaseLiabilityPaymentsDueYearTwo"`
			LesseeOperatingLeaseLiabilityPaymentsRemainderOfFiscalYear struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"LesseeOperatingLeaseLiabilityPaymentsRemainderOfFiscalYear"`
			LesseeOperatingLeaseLiabilityUndiscountedExcessAmount struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"LesseeOperatingLeaseLiabilityUndiscountedExcessAmount"`
			Liabilities struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"Liabilities"`
			LiabilitiesAndStockholdersEquity struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"LiabilitiesAndStockholdersEquity"`
			LiabilitiesCurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"LiabilitiesCurrent"`
			LiabilitiesNoncurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"LiabilitiesNoncurrent"`
			LongTermDebt struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"LongTermDebt"`
			LongTermDebtCurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"LongTermDebtCurrent"`
			LongTermDebtMaturitiesRepaymentsOfPrincipalAfterYearFive struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"LongTermDebtMaturitiesRepaymentsOfPrincipalAfterYearFive"`
			LongTermDebtMaturitiesRepaymentsOfPrincipalInNextTwelveMonths struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"LongTermDebtMaturitiesRepaymentsOfPrincipalInNextTwelveMonths"`
			LongTermDebtMaturitiesRepaymentsOfPrincipalInYearFive struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"LongTermDebtMaturitiesRepaymentsOfPrincipalInYearFive"`
			LongTermDebtMaturitiesRepaymentsOfPrincipalInYearFour struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"LongTermDebtMaturitiesRepaymentsOfPrincipalInYearFour"`
			LongTermDebtMaturitiesRepaymentsOfPrincipalInYearThree struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"LongTermDebtMaturitiesRepaymentsOfPrincipalInYearThree"`
			LongTermDebtMaturitiesRepaymentsOfPrincipalInYearTwo struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"LongTermDebtMaturitiesRepaymentsOfPrincipalInYearTwo"`
			LongTermDebtMaturitiesRepaymentsOfPrincipalRemainderOfFiscalYear struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"LongTermDebtMaturitiesRepaymentsOfPrincipalRemainderOfFiscalYear"`
			LongTermDebtNoncurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"LongTermDebtNoncurrent"`
			LongTermPurchaseCommitmentAmount struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"LongTermPurchaseCommitmentAmount"`
			MarketableSecuritiesCurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"MarketableSecuritiesCurrent"`
			MarketableSecuritiesNoncurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"MarketableSecuritiesNoncurrent"`
			MarketableSecuritiesRealizedGainLoss struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"MarketableSecuritiesRealizedGainLoss"`
			MarketingExpense struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"MarketingExpense"`
			NetCashProvidedByUsedInFinancingActivities struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"NetCashProvidedByUsedInFinancingActivities"`
			NetCashProvidedByUsedInFinancingActivitiesContinuingOperations struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"NetCashProvidedByUsedInFinancingActivitiesContinuingOperations"`
			NetCashProvidedByUsedInInvestingActivities struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"NetCashProvidedByUsedInInvestingActivities"`
			NetCashProvidedByUsedInInvestingActivitiesContinuingOperations struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"NetCashProvidedByUsedInInvestingActivitiesContinuingOperations"`
			NetCashProvidedByUsedInOperatingActivities struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"NetCashProvidedByUsedInOperatingActivities"`
			NetCashProvidedByUsedInOperatingActivitiesContinuingOperations struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"NetCashProvidedByUsedInOperatingActivitiesContinuingOperations"`
			NetIncomeLoss struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"NetIncomeLoss"`
			NoncurrentAssets struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"NoncurrentAssets"`
			NonoperatingIncomeExpense struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"NonoperatingIncomeExpense"`
			NontradeReceivablesCurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"NontradeReceivablesCurrent"`
			NotionalAmountOfForeignCurrencyDerivativeInstrumentsNotDesignatedAsHedgingInstruments struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"NotionalAmountOfForeignCurrencyDerivativeInstrumentsNotDesignatedAsHedgingInstruments"`
			NumberOfStores struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Store []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"Store"`
				} `json:"units"`
			} `json:"NumberOfStores"`
			OciBeforeReclassificationsBeforeTaxAttributableToParent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OciBeforeReclassificationsBeforeTaxAttributableToParent"`
			OperatingExpenses struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OperatingExpenses"`
			OperatingIncomeLoss struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OperatingIncomeLoss"`
			OperatingLeaseCost struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OperatingLeaseCost"`
			OperatingLeaseLiability struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OperatingLeaseLiability"`
			OperatingLeaseLiabilityCurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OperatingLeaseLiabilityCurrent"`
			OperatingLeaseLiabilityNoncurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OperatingLeaseLiabilityNoncurrent"`
			OperatingLeasePayments struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OperatingLeasePayments"`
			OperatingLeaseRightOfUseAsset struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OperatingLeaseRightOfUseAsset"`
			OperatingLeasesFutureMinimumPaymentsDue struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OperatingLeasesFutureMinimumPaymentsDue"`
			OperatingLeasesFutureMinimumPaymentsDueCurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OperatingLeasesFutureMinimumPaymentsDueCurrent"`
			OperatingLeasesFutureMinimumPaymentsDueInFiveYears struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OperatingLeasesFutureMinimumPaymentsDueInFiveYears"`
			OperatingLeasesFutureMinimumPaymentsDueInFourYears struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OperatingLeasesFutureMinimumPaymentsDueInFourYears"`
			OperatingLeasesFutureMinimumPaymentsDueInThreeYears struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OperatingLeasesFutureMinimumPaymentsDueInThreeYears"`
			OperatingLeasesFutureMinimumPaymentsDueInTwoYears struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OperatingLeasesFutureMinimumPaymentsDueInTwoYears"`
			OperatingLeasesFutureMinimumPaymentsDueThereafter struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OperatingLeasesFutureMinimumPaymentsDueThereafter"`
			OperatingLeasesRentExpenseNet struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OperatingLeasesRentExpenseNet"`
			OtherAccruedLiabilitiesCurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OtherAccruedLiabilitiesCurrent"`
			OtherAccruedLiabilitiesNoncurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OtherAccruedLiabilitiesNoncurrent"`
			OtherAssetsCurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OtherAssetsCurrent"`
			OtherAssetsNoncurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OtherAssetsNoncurrent"`
			OtherComprehensiveIncomeAvailableForSaleSecuritiesAdjustmentNetOfTaxPeriodIncreaseDecrease struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OtherComprehensiveIncomeAvailableForSaleSecuritiesAdjustmentNetOfTaxPeriodIncreaseDecrease"`
			OtherComprehensiveIncomeDerivativesQualifyingAsHedgesNetOfTaxPeriodIncreaseDecrease struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OtherComprehensiveIncomeDerivativesQualifyingAsHedgesNetOfTaxPeriodIncreaseDecrease"`
			OtherComprehensiveIncomeForeignCurrencyTransactionAndTranslationAdjustmentNetOfTaxPeriodIncreaseDecrease struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OtherComprehensiveIncomeForeignCurrencyTransactionAndTranslationAdjustmentNetOfTaxPeriodIncreaseDecrease"`
			OtherComprehensiveIncomeLossAvailableForSaleSecuritiesAdjustmentNetOfTax struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OtherComprehensiveIncomeLossAvailableForSaleSecuritiesAdjustmentNetOfTax"`
			OtherComprehensiveIncomeLossBeforeReclassificationsBeforeTax struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OtherComprehensiveIncomeLossBeforeReclassificationsBeforeTax"`
			OtherComprehensiveIncomeLossBeforeReclassificationsNetOfTax struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OtherComprehensiveIncomeLossBeforeReclassificationsNetOfTax"`
			OtherComprehensiveIncomeLossCashFlowHedgeGainLossBeforeReclassificationAndTax struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OtherComprehensiveIncomeLossCashFlowHedgeGainLossBeforeReclassificationAndTax"`
			OtherComprehensiveIncomeLossCashFlowHedgeGainLossReclassificationAfterTax struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OtherComprehensiveIncomeLossCashFlowHedgeGainLossReclassificationAfterTax"`
			OtherComprehensiveIncomeLossCashFlowHedgeGainLossReclassificationBeforeTax struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OtherComprehensiveIncomeLossCashFlowHedgeGainLossReclassificationBeforeTax"`
			OtherComprehensiveIncomeLossDerivativeExcludedComponentIncreaseDecreaseBeforeAdjustmentsAndTax struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OtherComprehensiveIncomeLossDerivativeExcludedComponentIncreaseDecreaseBeforeAdjustmentsAndTax"`
			OtherComprehensiveIncomeLossDerivativesQualifyingAsHedgesNetOfTax struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OtherComprehensiveIncomeLossDerivativesQualifyingAsHedgesNetOfTax"`
			OtherComprehensiveIncomeLossForeignCurrencyTransactionAndTranslationAdjustmentNetOfTax struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OtherComprehensiveIncomeLossForeignCurrencyTransactionAndTranslationAdjustmentNetOfTax"`
			OtherComprehensiveIncomeLossForeignCurrencyTranslationAdjustmentTax struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OtherComprehensiveIncomeLossForeignCurrencyTranslationAdjustmentTax"`
			OtherComprehensiveIncomeLossNetOfTax struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OtherComprehensiveIncomeLossNetOfTax"`
			OtherComprehensiveIncomeLossNetOfTaxPortionAttributableToParent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OtherComprehensiveIncomeLossNetOfTaxPortionAttributableToParent"`
			OtherComprehensiveIncomeLossReclassificationAdjustmentForSaleOfSecuritiesIncludedInNetIncomeNetOfTax struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OtherComprehensiveIncomeLossReclassificationAdjustmentForSaleOfSecuritiesIncludedInNetIncomeNetOfTax"`
			OtherComprehensiveIncomeLossReclassificationAdjustmentFromAOCIForSaleOfSecuritiesBeforeTax struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OtherComprehensiveIncomeLossReclassificationAdjustmentFromAOCIForSaleOfSecuritiesBeforeTax"`
			OtherComprehensiveIncomeLossReclassificationAdjustmentFromAOCIForSaleOfSecuritiesNetOfTax struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OtherComprehensiveIncomeLossReclassificationAdjustmentFromAOCIForSaleOfSecuritiesNetOfTax"`
			OtherComprehensiveIncomeLossReclassificationAdjustmentFromAOCIForSaleOfSecuritiesTax struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OtherComprehensiveIncomeLossReclassificationAdjustmentFromAOCIForSaleOfSecuritiesTax"`
			OtherComprehensiveIncomeLossReclassificationAdjustmentFromAOCIOnDerivativesBeforeTax struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OtherComprehensiveIncomeLossReclassificationAdjustmentFromAOCIOnDerivativesBeforeTax"`
			OtherComprehensiveIncomeLossReclassificationAdjustmentFromAOCIOnDerivativesNetOfTax struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OtherComprehensiveIncomeLossReclassificationAdjustmentFromAOCIOnDerivativesNetOfTax"`
			OtherComprehensiveIncomeLossReclassificationAdjustmentFromAOCIOnDerivativesTax struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OtherComprehensiveIncomeLossReclassificationAdjustmentFromAOCIOnDerivativesTax"`
			OtherComprehensiveIncomeLossReclassificationAdjustmentOnDerivativesIncludedInNetIncomeNetOfTax struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OtherComprehensiveIncomeLossReclassificationAdjustmentOnDerivativesIncludedInNetIncomeNetOfTax"`
			OtherComprehensiveIncomeLossReclassificationAdjustmentOnDerivativesIncludedInNetIncomeTax struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OtherComprehensiveIncomeLossReclassificationAdjustmentOnDerivativesIncludedInNetIncomeTax"`
			OtherComprehensiveIncomeLossTax struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OtherComprehensiveIncomeLossTax"`
			OtherComprehensiveIncomeLossTaxPortionAttributableToParent1 struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OtherComprehensiveIncomeLossTaxPortionAttributableToParent1"`
			OtherComprehensiveIncomeReclassificationAdjustmentOnDerivativesIncludedInNetIncomeNetOfTax struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OtherComprehensiveIncomeReclassificationAdjustmentOnDerivativesIncludedInNetIncomeNetOfTax"`
			OtherComprehensiveIncomeReclassificationAdjustmentOnDerivativesIncludedInNetIncomeTax struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OtherComprehensiveIncomeReclassificationAdjustmentOnDerivativesIncludedInNetIncomeTax"`
			OtherComprehensiveIncomeUnrealizedGainLossOnDerivativesArisingDuringPeriodNetOfTax struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OtherComprehensiveIncomeUnrealizedGainLossOnDerivativesArisingDuringPeriodNetOfTax"`
			OtherComprehensiveIncomeUnrealizedGainLossOnDerivativesArisingDuringPeriodTax struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OtherComprehensiveIncomeUnrealizedGainLossOnDerivativesArisingDuringPeriodTax"`
			OtherComprehensiveIncomeUnrealizedHoldingGainLossOnSecuritiesArisingDuringPeriodNetOfTax struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OtherComprehensiveIncomeUnrealizedHoldingGainLossOnSecuritiesArisingDuringPeriodNetOfTax"`
			OtherComprehensiveIncomeUnrealizedHoldingGainLossOnSecuritiesArisingDuringPeriodTax struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OtherComprehensiveIncomeUnrealizedHoldingGainLossOnSecuritiesArisingDuringPeriodTax"`
			OtherDeferredCreditsCurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OtherDeferredCreditsCurrent"`
			OtherLiabilitiesCurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OtherLiabilitiesCurrent"`
			OtherLiabilitiesNoncurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OtherLiabilitiesNoncurrent"`
			OtherNoncashIncomeExpense struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OtherNoncashIncomeExpense"`
			OtherNonoperatingIncomeExpense struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OtherNonoperatingIncomeExpense"`
			OtherPrepaidExpenseCurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OtherPrepaidExpenseCurrent"`
			OtherShortTermBorrowings struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"OtherShortTermBorrowings"`
			PaymentsForProceedsFromOtherInvestingActivities struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"PaymentsForProceedsFromOtherInvestingActivities"`
			PaymentsForRepurchaseOfCommonStock struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"PaymentsForRepurchaseOfCommonStock"`
			PaymentsOfDividends struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"PaymentsOfDividends"`
			PaymentsOfDividendsCommonStock struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"PaymentsOfDividendsCommonStock"`
			PaymentsRelatedToTaxWithholdingForShareBasedCompensation struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"PaymentsRelatedToTaxWithholdingForShareBasedCompensation"`
			PaymentsToAcquireAvailableForSaleSecurities struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"PaymentsToAcquireAvailableForSaleSecurities"`
			PaymentsToAcquireAvailableForSaleSecuritiesDebt struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"PaymentsToAcquireAvailableForSaleSecuritiesDebt"`
			PaymentsToAcquireBusinessesNetOfCashAcquired struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"PaymentsToAcquireBusinessesNetOfCashAcquired"`
			PaymentsToAcquireIntangibleAssets struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"PaymentsToAcquireIntangibleAssets"`
			PaymentsToAcquireOtherInvestments struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"PaymentsToAcquireOtherInvestments"`
			PaymentsToAcquireProductiveAssets struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"PaymentsToAcquireProductiveAssets"`
			PaymentsToAcquirePropertyPlantAndEquipment struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"PaymentsToAcquirePropertyPlantAndEquipment"`
			PreferredStockSharesAuthorized struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Shares []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"shares"`
				} `json:"units"`
			} `json:"PreferredStockSharesAuthorized"`
			PrepaidExpenseOtherNoncurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"PrepaidExpenseOtherNoncurrent"`
			ProceedsFromIssuanceOfCommonStock struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ProceedsFromIssuanceOfCommonStock"`
			ProceedsFromIssuanceOfLongTermDebt struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ProceedsFromIssuanceOfLongTermDebt"`
			ProceedsFromMaturitiesPrepaymentsAndCallsOfAvailableForSaleSecurities struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ProceedsFromMaturitiesPrepaymentsAndCallsOfAvailableForSaleSecurities"`
			ProceedsFromOtherShortTermDebt struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ProceedsFromOtherShortTermDebt"`
			ProceedsFromPaymentsForOtherFinancingActivities struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ProceedsFromPaymentsForOtherFinancingActivities"`
			ProceedsFromRepaymentsOfCommercialPaper struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ProceedsFromRepaymentsOfCommercialPaper"`
			ProceedsFromRepaymentsOfShortTermDebt struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ProceedsFromRepaymentsOfShortTermDebt"`
			ProceedsFromRepaymentsOfShortTermDebtMaturingInMoreThanThreeMonths struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ProceedsFromRepaymentsOfShortTermDebtMaturingInMoreThanThreeMonths"`
			ProceedsFromRepaymentsOfShortTermDebtMaturingInThreeMonthsOrLess struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ProceedsFromRepaymentsOfShortTermDebtMaturingInThreeMonthsOrLess"`
			ProceedsFromSaleAndMaturityOfOtherInvestments struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ProceedsFromSaleAndMaturityOfOtherInvestments"`
			ProceedsFromSaleOfAvailableForSaleSecurities struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ProceedsFromSaleOfAvailableForSaleSecurities"`
			ProceedsFromSaleOfAvailableForSaleSecuritiesDebt struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ProceedsFromSaleOfAvailableForSaleSecuritiesDebt"`
			ProceedsFromShortTermDebtMaturingInMoreThanThreeMonths struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ProceedsFromShortTermDebtMaturingInMoreThanThreeMonths"`
			ProductWarrantyAccrualClassifiedCurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ProductWarrantyAccrualClassifiedCurrent"`
			ProductWarrantyAccrualPayments struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ProductWarrantyAccrualPayments"`
			ProductWarrantyAccrualWarrantiesIssued struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ProductWarrantyAccrualWarrantiesIssued"`
			ProductWarrantyExpense struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ProductWarrantyExpense"`
			PropertyPlantAndEquipmentGross struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"PropertyPlantAndEquipmentGross"`
			PropertyPlantAndEquipmentNet struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"PropertyPlantAndEquipmentNet"`
			ReclassificationFromAccumulatedOtherComprehensiveIncomeCurrentPeriodBeforeTax struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ReclassificationFromAccumulatedOtherComprehensiveIncomeCurrentPeriodBeforeTax"`
			ReclassificationFromAccumulatedOtherComprehensiveIncomeCurrentPeriodNetOfTax struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ReclassificationFromAccumulatedOtherComprehensiveIncomeCurrentPeriodNetOfTax"`
			ReclassificationFromAociCurrentPeriodBeforeTaxAttributableToParent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ReclassificationFromAociCurrentPeriodBeforeTaxAttributableToParent"`
			RepaymentsOfLongTermDebt struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"RepaymentsOfLongTermDebt"`
			RepaymentsOfOtherShortTermDebt struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"RepaymentsOfOtherShortTermDebt"`
			RepaymentsOfShortTermDebtMaturingInMoreThanThreeMonths struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"RepaymentsOfShortTermDebtMaturingInMoreThanThreeMonths"`
			ResearchAndDevelopmentExpense struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ResearchAndDevelopmentExpense"`
			RestrictedCash struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"RestrictedCash"`
			RestrictedCashAndCashEquivalentsAtCarryingValue struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"RestrictedCashAndCashEquivalentsAtCarryingValue"`
			RestrictedCashAndCashEquivalentsNoncurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"RestrictedCashAndCashEquivalentsNoncurrent"`
			RestrictedInvestments struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"RestrictedInvestments"`
			RetainedEarningsAccumulatedDeficit struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"RetainedEarningsAccumulatedDeficit"`
			RevenueFromContractWithCustomerExcludingAssessedTax struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"RevenueFromContractWithCustomerExcludingAssessedTax"`
			Revenues struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"Revenues"`
			SalesRevenueNet struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"SalesRevenueNet"`
			SalesRevenueServicesGross struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"SalesRevenueServicesGross"`
			SecuritiesSoldUnderAgreementsToRepurchaseFairValueOfCollateral struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"SecuritiesSoldUnderAgreementsToRepurchaseFairValueOfCollateral"`
			SegmentReportingInformationOperatingIncomeLoss struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"SegmentReportingInformationOperatingIncomeLoss"`
			SegmentReportingReconcilingItemsAssets struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"SegmentReportingReconcilingItemsAssets"`
			SegmentReportingSegmentAssets struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"SegmentReportingSegmentAssets"`
			SellingGeneralAndAdministrativeExpense struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"SellingGeneralAndAdministrativeExpense"`
			ShareBasedCompensation struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ShareBasedCompensation"`
			ShareBasedCompensationArrangementByShareBasedPaymentAwardEquityInstrumentsOtherThanOptionsForfeitedInPeriod struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Shares []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"shares"`
				} `json:"units"`
			} `json:"ShareBasedCompensationArrangementByShareBasedPaymentAwardEquityInstrumentsOtherThanOptionsForfeitedInPeriod"`
			ShareBasedCompensationArrangementByShareBasedPaymentAwardEquityInstrumentsOtherThanOptionsForfeitedInPeriodWeightedAverageGrantDateFairValue struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ShareBasedCompensationArrangementByShareBasedPaymentAwardEquityInstrumentsOtherThanOptionsForfeitedInPeriodWeightedAverageGrantDateFairValue"`
			ShareBasedCompensationArrangementByShareBasedPaymentAwardEquityInstrumentsOtherThanOptionsForfeituresWeightedAverageGrantDateFairValue struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					USDShares []struct {
						Start string  `json:"start"`
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame,omitempty"`
					} `json:"USD/shares"`
				} `json:"units"`
			} `json:"ShareBasedCompensationArrangementByShareBasedPaymentAwardEquityInstrumentsOtherThanOptionsForfeituresWeightedAverageGrantDateFairValue"`
			ShareBasedCompensationArrangementByShareBasedPaymentAwardEquityInstrumentsOtherThanOptionsGrantsInPeriod struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Shares []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"shares"`
				} `json:"units"`
			} `json:"ShareBasedCompensationArrangementByShareBasedPaymentAwardEquityInstrumentsOtherThanOptionsGrantsInPeriod"`
			ShareBasedCompensationArrangementByShareBasedPaymentAwardEquityInstrumentsOtherThanOptionsGrantsInPeriodWeightedAverageGrantDateFairValue struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					USDShares []struct {
						Start string  `json:"start"`
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame,omitempty"`
					} `json:"USD/shares"`
				} `json:"units"`
			} `json:"ShareBasedCompensationArrangementByShareBasedPaymentAwardEquityInstrumentsOtherThanOptionsGrantsInPeriodWeightedAverageGrantDateFairValue"`
			ShareBasedCompensationArrangementByShareBasedPaymentAwardEquityInstrumentsOtherThanOptionsNonvestedIntrinsicValue struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					USDShares []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD/shares"`
				} `json:"units"`
			} `json:"ShareBasedCompensationArrangementByShareBasedPaymentAwardEquityInstrumentsOtherThanOptionsNonvestedIntrinsicValue"`
			ShareBasedCompensationArrangementByShareBasedPaymentAwardEquityInstrumentsOtherThanOptionsNonvestedNumber struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Shares []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"shares"`
				} `json:"units"`
			} `json:"ShareBasedCompensationArrangementByShareBasedPaymentAwardEquityInstrumentsOtherThanOptionsNonvestedNumber"`
			ShareBasedCompensationArrangementByShareBasedPaymentAwardEquityInstrumentsOtherThanOptionsNonvestedWeightedAverageGrantDateFairValue struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					USDShares []struct {
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame,omitempty"`
					} `json:"USD/shares"`
				} `json:"units"`
			} `json:"ShareBasedCompensationArrangementByShareBasedPaymentAwardEquityInstrumentsOtherThanOptionsNonvestedWeightedAverageGrantDateFairValue"`
			ShareBasedCompensationArrangementByShareBasedPaymentAwardEquityInstrumentsOtherThanOptionsVestedInPeriod struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Shares []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"shares"`
				} `json:"units"`
			} `json:"ShareBasedCompensationArrangementByShareBasedPaymentAwardEquityInstrumentsOtherThanOptionsVestedInPeriod"`
			ShareBasedCompensationArrangementByShareBasedPaymentAwardEquityInstrumentsOtherThanOptionsVestedInPeriodTotalFairValue struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ShareBasedCompensationArrangementByShareBasedPaymentAwardEquityInstrumentsOtherThanOptionsVestedInPeriodTotalFairValue"`
			ShareBasedCompensationArrangementByShareBasedPaymentAwardEquityInstrumentsOtherThanOptionsVestedInPeriodWeightedAverageGrantDateFairValue struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					USDShares []struct {
						Start string  `json:"start"`
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame,omitempty"`
					} `json:"USD/shares"`
				} `json:"units"`
			} `json:"ShareBasedCompensationArrangementByShareBasedPaymentAwardEquityInstrumentsOtherThanOptionsVestedInPeriodWeightedAverageGrantDateFairValue"`
			ShareBasedCompensationArrangementByShareBasedPaymentAwardFairValueAssumptionsExpectedDividendRate struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Pure []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"pure"`
				} `json:"units"`
			} `json:"ShareBasedCompensationArrangementByShareBasedPaymentAwardFairValueAssumptionsExpectedDividendRate"`
			ShareBasedCompensationArrangementByShareBasedPaymentAwardFairValueAssumptionsExpectedTerm struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Pure []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"pure"`
				} `json:"units"`
			} `json:"ShareBasedCompensationArrangementByShareBasedPaymentAwardFairValueAssumptionsExpectedTerm"`
			ShareBasedCompensationArrangementByShareBasedPaymentAwardFairValueAssumptionsRiskFreeInterestRate struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Pure []struct {
						Start string  `json:"start"`
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame"`
					} `json:"pure"`
				} `json:"units"`
			} `json:"ShareBasedCompensationArrangementByShareBasedPaymentAwardFairValueAssumptionsRiskFreeInterestRate"`
			ShareBasedCompensationArrangementByShareBasedPaymentAwardFairValueAssumptionsWeightedAverageVolatilityRate struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Pure []struct {
						Start string  `json:"start"`
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame"`
					} `json:"pure"`
				} `json:"units"`
			} `json:"ShareBasedCompensationArrangementByShareBasedPaymentAwardFairValueAssumptionsWeightedAverageVolatilityRate"`
			ShareBasedCompensationArrangementByShareBasedPaymentAwardNumberOfSharesAvailableForGrant struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Shares []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"shares"`
				} `json:"units"`
			} `json:"ShareBasedCompensationArrangementByShareBasedPaymentAwardNumberOfSharesAvailableForGrant"`
			ShareBasedCompensationArrangementByShareBasedPaymentAwardOptionsExercisableIntrinsicValue struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ShareBasedCompensationArrangementByShareBasedPaymentAwardOptionsExercisableIntrinsicValue"`
			ShareBasedCompensationArrangementByShareBasedPaymentAwardOptionsExercisableNumber struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Shares []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"shares"`
				} `json:"units"`
			} `json:"ShareBasedCompensationArrangementByShareBasedPaymentAwardOptionsExercisableNumber"`
			ShareBasedCompensationArrangementByShareBasedPaymentAwardOptionsExercisableWeightedAverageExercisePrice struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					USDShares []struct {
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame,omitempty"`
					} `json:"USD/shares"`
				} `json:"units"`
			} `json:"ShareBasedCompensationArrangementByShareBasedPaymentAwardOptionsExercisableWeightedAverageExercisePrice"`
			ShareBasedCompensationArrangementByShareBasedPaymentAwardOptionsExercisableWeightedAverageRemainingContractualTerm struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Pure []struct {
						Start string  `json:"start"`
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame,omitempty"`
					} `json:"pure"`
					Year []struct {
						Start string  `json:"start"`
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame,omitempty"`
					} `json:"Year"`
				} `json:"units"`
			} `json:"ShareBasedCompensationArrangementByShareBasedPaymentAwardOptionsExercisableWeightedAverageRemainingContractualTerm"`
			ShareBasedCompensationArrangementByShareBasedPaymentAwardOptionsExercisesInPeriod struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Shares []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"shares"`
				} `json:"units"`
			} `json:"ShareBasedCompensationArrangementByShareBasedPaymentAwardOptionsExercisesInPeriod"`
			ShareBasedCompensationArrangementByShareBasedPaymentAwardOptionsExercisesInPeriodTotalIntrinsicValue struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ShareBasedCompensationArrangementByShareBasedPaymentAwardOptionsExercisesInPeriodTotalIntrinsicValue"`
			ShareBasedCompensationArrangementByShareBasedPaymentAwardOptionsExercisesInPeriodWeightedAverageExercisePrice struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ShareBasedCompensationArrangementByShareBasedPaymentAwardOptionsExercisesInPeriodWeightedAverageExercisePrice"`
			ShareBasedCompensationArrangementByShareBasedPaymentAwardOptionsForfeituresInPeriod struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Shares []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"shares"`
				} `json:"units"`
			} `json:"ShareBasedCompensationArrangementByShareBasedPaymentAwardOptionsForfeituresInPeriod"`
			ShareBasedCompensationArrangementByShareBasedPaymentAwardOptionsForfeituresInPeriodWeightedAverageExercisePrice struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ShareBasedCompensationArrangementByShareBasedPaymentAwardOptionsForfeituresInPeriodWeightedAverageExercisePrice"`
			ShareBasedCompensationArrangementByShareBasedPaymentAwardOptionsGrantsInPeriod struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Shares []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"shares"`
				} `json:"units"`
			} `json:"ShareBasedCompensationArrangementByShareBasedPaymentAwardOptionsGrantsInPeriod"`
			ShareBasedCompensationArrangementByShareBasedPaymentAwardOptionsGrantsInPeriodWeightedAverageExercisePrice struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ShareBasedCompensationArrangementByShareBasedPaymentAwardOptionsGrantsInPeriodWeightedAverageExercisePrice"`
			ShareBasedCompensationArrangementByShareBasedPaymentAwardOptionsGrantsInPeriodWeightedAverageGrantDateFairValue struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					USDShares []struct {
						Start string  `json:"start"`
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame,omitempty"`
					} `json:"USD/shares"`
				} `json:"units"`
			} `json:"ShareBasedCompensationArrangementByShareBasedPaymentAwardOptionsGrantsInPeriodWeightedAverageGrantDateFairValue"`
			ShareBasedCompensationArrangementByShareBasedPaymentAwardOptionsOtherIncreasesDecreasesInPeriod struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Shares []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"shares"`
				} `json:"units"`
			} `json:"ShareBasedCompensationArrangementByShareBasedPaymentAwardOptionsOtherIncreasesDecreasesInPeriod"`
			ShareBasedCompensationArrangementByShareBasedPaymentAwardOptionsOutstandingIntrinsicValue struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"ShareBasedCompensationArrangementByShareBasedPaymentAwardOptionsOutstandingIntrinsicValue"`
			ShareBasedCompensationArrangementByShareBasedPaymentAwardOptionsOutstandingNumber struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Shares []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"shares"`
				} `json:"units"`
			} `json:"ShareBasedCompensationArrangementByShareBasedPaymentAwardOptionsOutstandingNumber"`
			ShareBasedCompensationArrangementByShareBasedPaymentAwardOptionsOutstandingWeightedAverageExercisePrice struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					USDShares []struct {
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame,omitempty"`
					} `json:"USD/shares"`
				} `json:"units"`
			} `json:"ShareBasedCompensationArrangementByShareBasedPaymentAwardOptionsOutstandingWeightedAverageExercisePrice"`
			ShareBasedCompensationArrangementByShareBasedPaymentAwardOptionsOutstandingWeightedAverageRemainingContractualTerm struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Pure []struct {
						Start string  `json:"start"`
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame,omitempty"`
					} `json:"pure"`
				} `json:"units"`
			} `json:"ShareBasedCompensationArrangementByShareBasedPaymentAwardOptionsOutstandingWeightedAverageRemainingContractualTerm"`
			ShareBasedCompensationArrangementByShareBasedPaymentAwardOptionsVestedAndExpectedToVestOutstandingNumber struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Shares []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"shares"`
				} `json:"units"`
			} `json:"ShareBasedCompensationArrangementByShareBasedPaymentAwardOptionsVestedAndExpectedToVestOutstandingNumber"`
			ShareBasedCompensationArrangementsByShareBasedPaymentAwardOptionsExercisesInPeriodWeightedAverageExercisePrice struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					USDShares []struct {
						Start string  `json:"start"`
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame,omitempty"`
					} `json:"USD/shares"`
				} `json:"units"`
			} `json:"ShareBasedCompensationArrangementsByShareBasedPaymentAwardOptionsExercisesInPeriodWeightedAverageExercisePrice"`
			ShareBasedCompensationArrangementsByShareBasedPaymentAwardOptionsForfeituresInPeriodWeightedAverageExercisePrice struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					USDShares []struct {
						Start string  `json:"start"`
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame,omitempty"`
					} `json:"USD/shares"`
				} `json:"units"`
			} `json:"ShareBasedCompensationArrangementsByShareBasedPaymentAwardOptionsForfeituresInPeriodWeightedAverageExercisePrice"`
			ShareBasedCompensationArrangementsByShareBasedPaymentAwardOptionsGrantsInPeriodWeightedAverageExercisePrice struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					USDShares []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD/shares"`
				} `json:"units"`
			} `json:"ShareBasedCompensationArrangementsByShareBasedPaymentAwardOptionsGrantsInPeriodWeightedAverageExercisePrice"`
			SharebasedCompensationArrangementBySharebasedPaymentAwardEquityInstrumentsOtherThanOptionsAggregateIntrinsicValueNonvested struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"SharebasedCompensationArrangementBySharebasedPaymentAwardEquityInstrumentsOtherThanOptionsAggregateIntrinsicValueNonvested"`
			SharebasedCompensationArrangementBySharebasedPaymentAwardOptionsExercisableIntrinsicValue1 struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"SharebasedCompensationArrangementBySharebasedPaymentAwardOptionsExercisableIntrinsicValue1"`
			SharebasedCompensationArrangementBySharebasedPaymentAwardOptionsOutstandingWeightedAverageRemainingContractualTerm1 struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Year []struct {
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame"`
					} `json:"Year"`
				} `json:"units"`
			} `json:"SharebasedCompensationArrangementBySharebasedPaymentAwardOptionsOutstandingWeightedAverageRemainingContractualTerm1"`
			SharesPaidForTaxWithholdingForShareBasedCompensation struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Shares []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"shares"`
				} `json:"units"`
			} `json:"SharesPaidForTaxWithholdingForShareBasedCompensation"`
			ShortTermDebtWeightedAverageInterestRate struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Pure []struct {
						End   string  `json:"end"`
						Val   float64 `json:"val"`
						Accn  string  `json:"accn"`
						Fy    int     `json:"fy"`
						Fp    string  `json:"fp"`
						Form  string  `json:"form"`
						Filed string  `json:"filed"`
						Frame string  `json:"frame,omitempty"`
					} `json:"pure"`
				} `json:"units"`
			} `json:"ShortTermDebtWeightedAverageInterestRate"`
			SignificantChangeInUnrecognizedTaxBenefitsIsReasonablyPossibleEstimatedRangeOfChangeLowerBound struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"SignificantChangeInUnrecognizedTaxBenefitsIsReasonablyPossibleEstimatedRangeOfChangeLowerBound"`
			SignificantChangeInUnrecognizedTaxBenefitsIsReasonablyPossibleEstimatedRangeOfChangeUpperBound struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"SignificantChangeInUnrecognizedTaxBenefitsIsReasonablyPossibleEstimatedRangeOfChangeUpperBound"`
			StandardProductWarrantyAccrual struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"StandardProductWarrantyAccrual"`
			StandardProductWarrantyAccrualPayments struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"StandardProductWarrantyAccrualPayments"`
			StandardProductWarrantyAccrualWarrantiesIssued struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"StandardProductWarrantyAccrualWarrantiesIssued"`
			StateAndLocalIncomeTaxExpenseBenefitContinuingOperations struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"StateAndLocalIncomeTaxExpenseBenefitContinuingOperations"`
			StockholdersEquity struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"StockholdersEquity"`
			StockholdersEquityNoteStockSplitConversionRatio1 struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Pure []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"pure"`
				} `json:"units"`
			} `json:"StockholdersEquityNoteStockSplitConversionRatio1"`
			StockIssuedDuringPeriodSharesStockOptionsExercised struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Shares []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"shares"`
				} `json:"units"`
			} `json:"StockIssuedDuringPeriodSharesStockOptionsExercised"`
			StockIssuedDuringPeriodValueAcquisitions struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"StockIssuedDuringPeriodValueAcquisitions"`
			StockIssuedDuringPeriodValueShareBasedCompensation struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"StockIssuedDuringPeriodValueShareBasedCompensation"`
			StockRepurchasedAndRetiredDuringPeriodShares struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Shares []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"shares"`
				} `json:"units"`
			} `json:"StockRepurchasedAndRetiredDuringPeriodShares"`
			StockRepurchasedAndRetiredDuringPeriodValue struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"StockRepurchasedAndRetiredDuringPeriodValue"`
			StockRepurchaseProgramAuthorizedAmount struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"StockRepurchaseProgramAuthorizedAmount"`
			StockRepurchaseProgramAuthorizedAmount1 struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"StockRepurchaseProgramAuthorizedAmount1"`
			TaxAdjustmentsSettlementsAndUnusualProvisions struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"TaxAdjustmentsSettlementsAndUnusualProvisions"`
			TaxesPayableCurrent struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"TaxesPayableCurrent"`
			TranslationAdjustmentForNetInvestmentHedgeNetOfTax struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"TranslationAdjustmentForNetInvestmentHedgeNetOfTax"`
			TreasuryStockValueAcquiredCostMethod struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"TreasuryStockValueAcquiredCostMethod"`
			UndistributedEarningsOfForeignSubsidiaries struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"UndistributedEarningsOfForeignSubsidiaries"`
			UnrecognizedTaxBenefits struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"UnrecognizedTaxBenefits"`
			UnrecognizedTaxBenefitsDecreasesResultingFromPriorPeriodTaxPositions struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"UnrecognizedTaxBenefitsDecreasesResultingFromPriorPeriodTaxPositions"`
			UnrecognizedTaxBenefitsDecreasesResultingFromSettlementsWithTaxingAuthorities struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"UnrecognizedTaxBenefitsDecreasesResultingFromSettlementsWithTaxingAuthorities"`
			UnrecognizedTaxBenefitsIncomeTaxPenaltiesAndInterestAccrued struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"UnrecognizedTaxBenefitsIncomeTaxPenaltiesAndInterestAccrued"`
			UnrecognizedTaxBenefitsIncomeTaxPenaltiesAndInterestExpense struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"UnrecognizedTaxBenefitsIncomeTaxPenaltiesAndInterestExpense"`
			UnrecognizedTaxBenefitsIncreasesResultingFromCurrentPeriodTaxPositions struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"UnrecognizedTaxBenefitsIncreasesResultingFromCurrentPeriodTaxPositions"`
			UnrecognizedTaxBenefitsIncreasesResultingFromPriorPeriodTaxPositions struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"UnrecognizedTaxBenefitsIncreasesResultingFromPriorPeriodTaxPositions"`
			UnrecognizedTaxBenefitsPeriodIncreaseDecrease struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"UnrecognizedTaxBenefitsPeriodIncreaseDecrease"`
			UnrecognizedTaxBenefitsReductionsResultingFromLapseOfApplicableStatuteOfLimitations struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"UnrecognizedTaxBenefitsReductionsResultingFromLapseOfApplicableStatuteOfLimitations"`
			UnrecognizedTaxBenefitsThatWouldImpactEffectiveTaxRate struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"UnrecognizedTaxBenefitsThatWouldImpactEffectiveTaxRate"`
			UnrecordedUnconditionalPurchaseObligationBalanceOnFifthAnniversary struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"UnrecordedUnconditionalPurchaseObligationBalanceOnFifthAnniversary"`
			UnrecordedUnconditionalPurchaseObligationBalanceOnFirstAnniversary struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"UnrecordedUnconditionalPurchaseObligationBalanceOnFirstAnniversary"`
			UnrecordedUnconditionalPurchaseObligationBalanceOnFourthAnniversary struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"UnrecordedUnconditionalPurchaseObligationBalanceOnFourthAnniversary"`
			UnrecordedUnconditionalPurchaseObligationBalanceOnSecondAnniversary struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"UnrecordedUnconditionalPurchaseObligationBalanceOnSecondAnniversary"`
			UnrecordedUnconditionalPurchaseObligationBalanceOnThirdAnniversary struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"UnrecordedUnconditionalPurchaseObligationBalanceOnThirdAnniversary"`
			UnrecordedUnconditionalPurchaseObligationBalanceSheetAmount struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"UnrecordedUnconditionalPurchaseObligationBalanceSheetAmount"`
			UnrecordedUnconditionalPurchaseObligationDueAfterFiveYears struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"UnrecordedUnconditionalPurchaseObligationDueAfterFiveYears"`
			VariableLeaseCost struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Usd []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int64  `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"USD"`
				} `json:"units"`
			} `json:"VariableLeaseCost"`
			WeightedAverageNumberDilutedSharesOutstandingAdjustment struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Shares []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"shares"`
				} `json:"units"`
			} `json:"WeightedAverageNumberDilutedSharesOutstandingAdjustment"`
			WeightedAverageNumberOfDilutedSharesOutstanding struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Shares []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"shares"`
				} `json:"units"`
			} `json:"WeightedAverageNumberOfDilutedSharesOutstanding"`
			WeightedAverageNumberOfSharesOutstandingBasic struct {
				Label       string `json:"label"`
				Description string `json:"description"`
				Units       struct {
					Shares []struct {
						Start string `json:"start"`
						End   string `json:"end"`
						Val   int    `json:"val"`
						Accn  string `json:"accn"`
						Fy    int    `json:"fy"`
						Fp    string `json:"fp"`
						Form  string `json:"form"`
						Filed string `json:"filed"`
						Frame string `json:"frame,omitempty"`
					} `json:"shares"`
				} `json:"units"`
			} `json:"WeightedAverageNumberOfSharesOutstandingBasic"`
		} `json:"us-gaap"`
	} `json:"facts"`
}