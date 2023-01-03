// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package v1

// Driven from e2ap_constants.proto
// TODO: Automate the generation of this file

type ProcedureCodeT int32

const (
	ProcedureCodeIDhandoverPreparation                                  ProcedureCodeT = 0
	ProcedureCodeIDsNStatusTransfer                                     ProcedureCodeT = 1
	ProcedureCodeIDhandoverCancel                                       ProcedureCodeT = 2
	ProcedureCodeIDretrieveUEContext                                    ProcedureCodeT = 3
	ProcedureCodeIDrANPaging                                            ProcedureCodeT = 4
	ProcedureCodeIDxnUAddressIndication                                 ProcedureCodeT = 5
	ProcedureCodeIDuEContextRelease                                     ProcedureCodeT = 6
	ProcedureCodeIDsNGRANnodeAdditionPreparation                        ProcedureCodeT = 7
	ProcedureCodeIDsNGRANnodeReconfigurationCompletion                  ProcedureCodeT = 8
	ProcedureCodeIDmNGRANnodeinitiatedSNGRANnodeModificationPreparation ProcedureCodeT = 9
	ProcedureCodeIDsNGRANnodeinitiatedSNGRANnodeModificationPreparation ProcedureCodeT = 10
	ProcedureCodeIDmNGRANnodeinitiatedSNGRANnodeRelease                 ProcedureCodeT = 11
	ProcedureCodeIDsNGRANnodeinitiatedSNGRANnodeRelease                 ProcedureCodeT = 12
	ProcedureCodeIDsNGRANnodeCounterCheck                               ProcedureCodeT = 13
	ProcedureCodeIDsNGRANnodeChange                                     ProcedureCodeT = 14
	ProcedureCodeIDrRCTransfer                                          ProcedureCodeT = 15
	ProcedureCodeIDxnRemoval                                            ProcedureCodeT = 16
	ProcedureCodeIDxnSetup                                              ProcedureCodeT = 17
	ProcedureCodeIDnGRANnodeConfigurationUpdate                         ProcedureCodeT = 18
	ProcedureCodeIDcellActivation                                       ProcedureCodeT = 19
	ProcedureCodeIDreset                                                ProcedureCodeT = 20
	ProcedureCodeIDerrorIndication                                      ProcedureCodeT = 21
	ProcedureCodeIDprivateMessage                                       ProcedureCodeT = 22
	ProcedureCodeIDnotificationControl                                  ProcedureCodeT = 23
	ProcedureCodeIDactivityNotification                                 ProcedureCodeT = 24
	ProcedureCodeIDEUTRANRCellResourceCoordination                      ProcedureCodeT = 25
	ProcedureCodeIDsecondaryRATDataUsageReport                          ProcedureCodeT = 26
	ProcedureCodeIDdeactivateTrace                                      ProcedureCodeT = 27
	ProcedureCodeIDtraceStart                                           ProcedureCodeT = 28
	ProcedureCodeIDhandoverSuccess                                      ProcedureCodeT = 29
	ProcedureCodeIDconditionalHandoverCancel                            ProcedureCodeT = 30
	ProcedureCodeIDearlyStatusTransfer                                  ProcedureCodeT = 31
	ProcedureCodeIDfailureIndication                                    ProcedureCodeT = 32
	ProcedureCodeIDhandoverReport                                       ProcedureCodeT = 33
	ProcedureCodeIDresourceStatusReportingInitiation                    ProcedureCodeT = 34
	ProcedureCodeIDresourceStatusReporting                              ProcedureCodeT = 35
	ProcedureCodeIDmobilitySettingsChange                               ProcedureCodeT = 36
	ProcedureCodeIDaccessAndMobilityIndication                          ProcedureCodeT = 37
)

type ProtocolIeID int32

const (
	ProtocolIeIDActivatedServedCells                         ProtocolIeID = 0
	ProtocolIeIDActivationIDforCellActivation                ProtocolIeID = 1
	ProtocolIeIDadmittedSplitSRB                             ProtocolIeID = 2
	ProtocolIeIDadmittedSplitSRBrelease                      ProtocolIeID = 3
	ProtocolIeIDAMFRegionInformation                         ProtocolIeID = 4
	ProtocolIeIDAssistanceDataForRANPaging                   ProtocolIeID = 5
	ProtocolIeIDBearersSubjectToCounterCheck                 ProtocolIeID = 6
	ProtocolIeIDCause                                        ProtocolIeID = 7
	ProtocolIeIDcellAssistanceInfoNR                         ProtocolIeID = 8
	ProtocolIeIDConfigurationUpdateInitiatingNodeChoice      ProtocolIeID = 9
	ProtocolIeIDCriticalityDiagnostics                       ProtocolIeID = 10
	ProtocolIeIDXnUAddressInfoperPDUSessionList              ProtocolIeID = 11
	ProtocolIeIDDRBsSubjectToStatusTransferList              ProtocolIeID = 12
	ProtocolIeIDExpectedUEBehaviour                          ProtocolIeID = 13
	ProtocolIeIDGlobalNGRANnodeID                            ProtocolIeID = 14
	ProtocolIeIDGUAMI                                        ProtocolIeID = 15
	ProtocolIeIDindexToRatFrequSelectionPriority             ProtocolIeID = 16
	ProtocolIeIDinitiatingNodeTypeResourceCoordRequest       ProtocolIeID = 17
	ProtocolIeIDListofservedcellsEUTRA                       ProtocolIeID = 18
	ProtocolIeIDListofservedcellsNR                          ProtocolIeID = 19
	ProtocolIeIDLocationReportingInformation                 ProtocolIeID = 20
	ProtocolIeIDMACI                                         ProtocolIeID = 21
	ProtocolIeIDMaskedIMEISV                                 ProtocolIeID = 22
	ProtocolIeIDMNGRANnodeUEXnAPID                           ProtocolIeID = 23
	ProtocolIeIDMNtoSNContainer                              ProtocolIeID = 24
	ProtocolIeIDMobilityRestrictionList                      ProtocolIeID = 25
	ProtocolIeIDnewNGRANCellIdentity                         ProtocolIeID = 26
	ProtocolIeIDnewNGRANnodeUEXnAPID                         ProtocolIeID = 27
	ProtocolIeIDUEReportRRCTransfer                          ProtocolIeID = 28
	ProtocolIeIDoldNGRANnodeUEXnAPID                         ProtocolIeID = 29
	ProtocolIeIDOldtoNewNGRANnodeResumeContainer             ProtocolIeID = 30
	ProtocolIeIDPagingDRX                                    ProtocolIeID = 31
	ProtocolIeIDPCellID                                      ProtocolIeID = 32
	ProtocolIeIDPDCPChangeIndication                         ProtocolIeID = 33
	ProtocolIeIDPDUSessionAdmittedAddedAddReqAck             ProtocolIeID = 34
	ProtocolIeIDPDUSessionAdmittedModSNModConfirm            ProtocolIeID = 35
	ProtocolIeIDPDUSessionAdmittedSNModResponse              ProtocolIeID = 36
	ProtocolIeIDPDUSessionNotAdmittedAddReqAck               ProtocolIeID = 37
	ProtocolIeIDPDUSessionNotAdmittedSNModResponse           ProtocolIeID = 38
	ProtocolIeIDPDUSessionReleasedListRelConf                ProtocolIeID = 39
	ProtocolIeIDPDUSessionReleasedSNModConfirm               ProtocolIeID = 40
	ProtocolIeIDPDUSessionResourcesActivityNotifyList        ProtocolIeID = 41
	ProtocolIeIDPDUSessionResourcesAdmittedList              ProtocolIeID = 42
	ProtocolIeIDPDUSessionResourcesNotAdmittedList           ProtocolIeID = 43
	ProtocolIeIDPDUSessionResourcesNotifyList                ProtocolIeID = 44
	ProtocolIeIDPDUSessionSNChangeConfirmList                ProtocolIeID = 45
	ProtocolIeIDPDUSessionSNChangeRequiredList               ProtocolIeID = 46
	ProtocolIeIDPDUSessionToBeAddedAddReq                    ProtocolIeID = 47
	ProtocolIeIDPDUSessionToBeModifiedSNModRequired          ProtocolIeID = 48
	ProtocolIeIDPDUSessionToBeReleasedListRelRqd             ProtocolIeID = 49
	ProtocolIeIDPDUSessionToBeReleasedRelReq                 ProtocolIeID = 50
	ProtocolIeIDPDUSessionToBeReleasedSNModRequired          ProtocolIeID = 51
	ProtocolIeIDRANPagingArea                                ProtocolIeID = 52
	ProtocolIeIDPagingPriority                               ProtocolIeID = 53
	ProtocolIeIDrequestedSplitSRB                            ProtocolIeID = 54
	ProtocolIeIDrequestedSplitSRBrelease                     ProtocolIeID = 55
	ProtocolIeIDResetRequestTypeInfo                         ProtocolIeID = 56
	ProtocolIeIDResetResponseTypeInfo                        ProtocolIeID = 57
	ProtocolIeIDRespondingNodeTypeConfigUpdateAck            ProtocolIeID = 58
	ProtocolIeIDrespondingNodeTypeResourceCoordResponse      ProtocolIeID = 59
	ProtocolIeIDResponseInfoReconfCompl                      ProtocolIeID = 60
	ProtocolIeIDRRCConfigIndication                          ProtocolIeID = 61
	ProtocolIeIDRRCResumeCause                               ProtocolIeID = 62
	ProtocolIeIDSCGConfigurationQuery                        ProtocolIeID = 63
	ProtocolIeIDselectedPLMN                                 ProtocolIeID = 64
	ProtocolIeIDServedCellsToActivate                        ProtocolIeID = 65
	ProtocolIeIDservedCellsToUpdateEUTRA                     ProtocolIeID = 66
	ProtocolIeIDServedCellsToUpdateInitiatingNodeChoice      ProtocolIeID = 67
	ProtocolIeIDservedCellsToUpdateNR                        ProtocolIeID = 68
	ProtocolIeIDsngRANnodeSecurityKey                        ProtocolIeID = 69
	ProtocolIeIDSNGRANnodeUEAMBR                             ProtocolIeID = 70
	ProtocolIeIDSNGRANnodeUEXnAPID                           ProtocolIeID = 71
	ProtocolIeIDSNtoMNContainer                              ProtocolIeID = 72
	ProtocolIeIDsourceNGRANnodeUEXnAPID                      ProtocolIeID = 73
	ProtocolIeIDSplitSRBRRCTransfer                          ProtocolIeID = 74
	ProtocolIeIDTAISupportlist                               ProtocolIeID = 75
	ProtocolIeIDTimeToWait                                   ProtocolIeID = 76
	ProtocolIeIDTarget2SourceNGRANnodeTranspContainer        ProtocolIeID = 77
	ProtocolIeIDtargetCellGlobalID                           ProtocolIeID = 78
	ProtocolIeIDtargetNGRANnodeUEXnAPID                      ProtocolIeID = 79
	ProtocolIeIDtargetSNGRANnodeID                           ProtocolIeID = 80
	ProtocolIeIDTraceActivation                              ProtocolIeID = 81
	ProtocolIeIDUEContextID                                  ProtocolIeID = 82
	ProtocolIeIDUEContextInfoHORequest                       ProtocolIeID = 83
	ProtocolIeIDUEContextInfoRetrUECtxtResp                  ProtocolIeID = 84
	ProtocolIeIDUEContextInfoSNModRequest                    ProtocolIeID = 85
	ProtocolIeIDUEContextKeptIndicator                       ProtocolIeID = 86
	ProtocolIeIDUEContextRefAtSNHORequest                    ProtocolIeID = 87
	ProtocolIeIDUEHistoryInformation                         ProtocolIeID = 88
	ProtocolIeIDUEIdentityIndexValue                         ProtocolIeID = 89
	ProtocolIeIDUERANPagingIdentity                          ProtocolIeID = 90
	ProtocolIeIDUESecurityCapabilities                       ProtocolIeID = 91
	ProtocolIeIDUserPlaneTrafficActivityReport               ProtocolIeID = 92
	ProtocolIeIDXnRemovalThreshold                           ProtocolIeID = 93
	ProtocolIeIDDesiredActNotificationLevel                  ProtocolIeID = 94
	ProtocolIeIDAvailableDRBIDs                              ProtocolIeID = 95
	ProtocolIeIDAdditionalDRBIDs                             ProtocolIeID = 96
	ProtocolIeIDSpareDRBIDs                                  ProtocolIeID = 97
	ProtocolIeIDRequiredNumberOfDRBIDs                       ProtocolIeID = 98
	ProtocolIeIDTNLAToAddList                                ProtocolIeID = 99
	ProtocolIeIDTNLAToUpdateList                             ProtocolIeID = 100
	ProtocolIeIDTNLAToRemoveList                             ProtocolIeID = 101
	ProtocolIeIDTNLASetupList                                ProtocolIeID = 102
	ProtocolIeIDTNLAFailedToSetupList                        ProtocolIeID = 103
	ProtocolIeIDPDUSessionToBeReleasedRelReqAck              ProtocolIeID = 104
	ProtocolIeIDSNGRANnodeMaxIPDataRateUL                    ProtocolIeID = 105
	ProtocolIeIDPDUSessionResourceSecondaryRATUsageList      ProtocolIeID = 107
	ProtocolIeIDAdditionalULNGUTNLatUPFList                  ProtocolIeID = 108
	ProtocolIeIDSecondarydataForwardingInfoFromTargetList    ProtocolIeID = 109
	ProtocolIeIDLocationInformationSNReporting               ProtocolIeID = 110
	ProtocolIeIDLocationInformationSN                        ProtocolIeID = 111
	ProtocolIeIDLastEUTRANPLMNIdentity                       ProtocolIeID = 112
	ProtocolIeIDSNGRANnodeMaxIPDataRateDL                    ProtocolIeID = 113
	ProtocolIeIDMaxIPrateDL                                  ProtocolIeID = 114
	ProtocolIeIDSecurityResult                               ProtocolIeID = 115
	ProtocolIeIDSNSSAI                                       ProtocolIeID = 116
	ProtocolIeIDMRDCResourceCoordinationInfo                 ProtocolIeID = 117
	ProtocolIeIDAMFRegionInformationToAdd                    ProtocolIeID = 118
	ProtocolIeIDAMFRegionInformationToDelete                 ProtocolIeID = 119
	ProtocolIeIDOldQoSFlowMapULendmarkerexpected             ProtocolIeID = 120
	ProtocolIeIDRANPagingFailure                             ProtocolIeID = 121
	ProtocolIeIDUERadioCapabilityForPaging                   ProtocolIeID = 122
	ProtocolIeIDPDUSessionDataForwardingSNModResponse        ProtocolIeID = 123
	ProtocolIeIDDRBsNotAdmittedSetupModifyList               ProtocolIeID = 124
	ProtocolIeIDSecondaryMNXnUTNLInfoatM                     ProtocolIeID = 125
	ProtocolIeIDNEDCTDMPattern                               ProtocolIeID = 126
	ProtocolIeIDPDUSessionCommonNetworkInstance              ProtocolIeID = 127
	ProtocolIeIDBPLMNIDInfoEUTRA                             ProtocolIeID = 128
	ProtocolIeIDBPLMNIDInfoNR                                ProtocolIeID = 129
	ProtocolIeIDInterfaceInstanceIndication                  ProtocolIeID = 130
	ProtocolIeIDSNGRANnodeAdditionTriggerInd                 ProtocolIeID = 131
	ProtocolIeIDDefaultDRBAllowed                            ProtocolIeID = 132
	ProtocolIeIDDRBIDstakenintouse                           ProtocolIeID = 133
	ProtocolIeIDSplitSessionIndicator                        ProtocolIeID = 134
	ProtocolIeIDCNTypeRestrictionsForEquivalent              ProtocolIeID = 135
	ProtocolIeIDCNTypeRestrictionsForServing                 ProtocolIeID = 136
	ProtocolIeIDDRBstransferredtoMN                          ProtocolIeID = 137
	ProtocolIeIDULForwardingProposal                         ProtocolIeID = 138
	ProtocolIeIDEndpointIPAddressAndPort                     ProtocolIeID = 139
	ProtocolIeIDIntendedTDDDLULConfigurationNR               ProtocolIeID = 140
	ProtocolIeIDTNLConfigurationInfo                         ProtocolIeID = 141
	ProtocolIeIDPartialListIndicatorNR                       ProtocolIeID = 142
	ProtocolIeIDMessageOversizeNotification                  ProtocolIeID = 143
	ProtocolIeIDCellAndCapacityAssistanceInfoNR              ProtocolIeID = 144
	ProtocolIeIDNGRANTraceID                                 ProtocolIeID = 145
	ProtocolIeIDNonGBRResourcesOffered                       ProtocolIeID = 146
	ProtocolIeIDFastMCGRecoveryRRCTransferSNtoMN             ProtocolIeID = 147
	ProtocolIeIDRequestedFastMCGRecoveryViaSRB3              ProtocolIeID = 148
	ProtocolIeIDAvailableFastMCGRecoveryViaSRB3              ProtocolIeID = 149
	ProtocolIeIDRequestedFastMCGRecoveryViaSRB3Release       ProtocolIeID = 150
	ProtocolIeIDReleaseFastMCGRecoveryViaSRB3                ProtocolIeID = 151
	ProtocolIeIDFastMCGRecoveryRRCTransferMNtoSN             ProtocolIeID = 152
	ProtocolIeIDExtendedRATRestrictionInformation            ProtocolIeID = 153
	ProtocolIeIDQoSMonitoringRequest                         ProtocolIeID = 154
	ProtocolIeIDFiveGCMobilityRestrictionListContainer       ProtocolIeID = 155
	ProtocolIeIDPartialListIndicatorEUTRA                    ProtocolIeID = 156
	ProtocolIeIDCellAndCapacityAssistanceInfoEUTRA           ProtocolIeID = 157
	ProtocolIeIDCHOinformationReq                            ProtocolIeID = 158
	ProtocolIeIDCHOinformationAck                            ProtocolIeID = 159
	ProtocolIeIDtargetCellsToCancel                          ProtocolIeID = 160
	ProtocolIeIDrequestedTargetCellGlobalID                  ProtocolIeID = 161
	ProtocolIeIDprocedureStage                               ProtocolIeID = 162
	ProtocolIeIDDAPSRequestInfo                              ProtocolIeID = 163
	ProtocolIeIDDAPSResponseInfoList                         ProtocolIeID = 164
	ProtocolIeIDCHOMRDCIndicator                             ProtocolIeID = 165
	ProtocolIeIDOffsetOfNbiotChannelNumberToDLEARFCN         ProtocolIeID = 166
	ProtocolIeIDOffsetOfNbiotChannelNumberToULEARFCN         ProtocolIeID = 167
	ProtocolIeIDNBIoTULDLAlignmentOffset                     ProtocolIeID = 168
	ProtocolIeIDLTEV2XServicesAuthorized                     ProtocolIeID = 169
	ProtocolIeIDNRV2XServicesAuthorized                      ProtocolIeID = 170
	ProtocolIeIDLTEUESidelinkAggregateMaximumBitRate         ProtocolIeID = 171
	ProtocolIeIDNRUESidelinkAggregateMaximumBitRate          ProtocolIeID = 172
	ProtocolIeIDPC5QoSParameters                             ProtocolIeID = 173
	ProtocolIeIDAlternativeQoSParaSetList                    ProtocolIeID = 174
	ProtocolIeIDCurrentQoSParaSetIndex                       ProtocolIeID = 175
	ProtocolIeIDMobilityInformation                          ProtocolIeID = 176
	ProtocolIeIDInitiatingConditionFailureIndication         ProtocolIeID = 177
	ProtocolIeIDUEHistoryInformationFromTheUE                ProtocolIeID = 178
	ProtocolIeIDHandoverReportType                           ProtocolIeID = 179
	ProtocolIeIDHandoverCause                                ProtocolIeID = 180
	ProtocolIeIDSourceCellCGI                                ProtocolIeID = 181
	ProtocolIeIDTargetCellCGI                                ProtocolIeID = 182
	ProtocolIeIDReEstablishmentCellCGI                       ProtocolIeID = 183
	ProtocolIeIDTargetCellinEUTRAN                           ProtocolIeID = 184
	ProtocolIeIDSourceCellCRNTI                              ProtocolIeID = 185
	ProtocolIeIDUERLFReportContainer                         ProtocolIeID = 186
	ProtocolIeIDNGRANNode1MeasurementID                      ProtocolIeID = 187
	ProtocolIeIDNGRANNode2MeasurementID                      ProtocolIeID = 188
	ProtocolIeIDRegistrationRequest                          ProtocolIeID = 189
	ProtocolIeIDReportCharacteristics                        ProtocolIeID = 190
	ProtocolIeIDCellToReport                                 ProtocolIeID = 191
	ProtocolIeIDReportingPeriodicity                         ProtocolIeID = 192
	ProtocolIeIDCellMeasurementResult                        ProtocolIeID = 193
	ProtocolIeIDNGRANnode1CellID                             ProtocolIeID = 194
	ProtocolIeIDNGRANnode2CellID                             ProtocolIeID = 195
	ProtocolIeIDNGRANnode1MobilityParameters                 ProtocolIeID = 196
	ProtocolIeIDNGRANnode2ProposedMobilityParameters         ProtocolIeID = 197
	ProtocolIeIDMobilityParametersModificationRange          ProtocolIeID = 198
	ProtocolIeIDTDDULDLConfigurationCommonNR                 ProtocolIeID = 199
	ProtocolIeIDCarrierList                                  ProtocolIeID = 200
	ProtocolIeIDULCarrierList                                ProtocolIeID = 201
	ProtocolIeIDFrequencyShift7p5khz                         ProtocolIeID = 202
	ProtocolIeIDSSBPositionsInBurst                          ProtocolIeID = 203
	ProtocolIeIDNRCellPRACHConfig                            ProtocolIeID = 204
	ProtocolIeIDRACHReportInformation                        ProtocolIeID = 205
	ProtocolIeIDIABNodeIndication                            ProtocolIeID = 206
	ProtocolIeIDRedundantULNGUTNLatUPF                       ProtocolIeID = 207
	ProtocolIeIDCNPacketDelayBudgetDownlink                  ProtocolIeID = 208
	ProtocolIeIDCNPacketDelayBudgetUplink                    ProtocolIeID = 209
	ProtocolIeIDAdditionalRedundantULNGUTNLatUPFList         ProtocolIeID = 210
	ProtocolIeIDRedundantCommonNetworkInstance               ProtocolIeID = 211
	ProtocolIeIDTSCTrafficCharacteristics                    ProtocolIeID = 212
	ProtocolIeIDRedundantQoSFlowIndicator                    ProtocolIeID = 213
	ProtocolIeIDRedundantDLNGUTNLatNGRAN                     ProtocolIeID = 214
	ProtocolIeIDExtendedPacketDelayBudget                    ProtocolIeID = 215
	ProtocolIeIDAdditionalPDCPDuplicationTNLList             ProtocolIeID = 216
	ProtocolIeIDRedundantPDUSessionInformation               ProtocolIeID = 217
	ProtocolIeIDUsedRSNInformation                           ProtocolIeID = 218
	ProtocolIeIDRLCDuplicationInformation                    ProtocolIeID = 219
	ProtocolIeIDNPNBroadcastInformation                      ProtocolIeID = 220
	ProtocolIeIDNPNPagingAssistanceInformation               ProtocolIeID = 221
	ProtocolIeIDNPNMobilityInformation                       ProtocolIeID = 222
	ProtocolIeIDNPNSupport                                   ProtocolIeID = 223
	ProtocolIeIDMDTConfiguration                             ProtocolIeID = 224
	ProtocolIeIDMDTPLMNList                                  ProtocolIeID = 225
	ProtocolIeIDTraceCollectionEntityURI                     ProtocolIeID = 226
	ProtocolIeIDUERadioCapabilityID                          ProtocolIeID = 227
	ProtocolIeIDCSIRSTransmissionIndication                  ProtocolIeID = 228
	ProtocolIeIDSNTriggered                                  ProtocolIeID = 229
	ProtocolIeIDDLCarrierList                                ProtocolIeID = 230
	ProtocolIeIDExtendedTAISliceSupportList                  ProtocolIeID = 231
	ProtocolIeIDcellAssistanceInfoEUTRA                      ProtocolIeID = 232
	ProtocolIeIDConfiguredTACIndication                      ProtocolIeID = 233
	ProtocolIeIDsecondarySNULPDCPUPTNLInfo                   ProtocolIeID = 234
	ProtocolIeIDpdcpDuplicationConfiguration                 ProtocolIeID = 235
	ProtocolIeIDduplicationActivation                        ProtocolIeID = 236
	ProtocolIeIDNPRACHConfiguration                          ProtocolIeID = 237
	ProtocolIeIDQosMonitoringReportingFrequency              ProtocolIeID = 238
	ProtocolIeIDQoSFlowsMappedtoDRBSetupResponseMNterminated ProtocolIeID = 239
	ProtocolIeIDDLschedulingPDCCHCCEusage                    ProtocolIeID = 240
	ProtocolIeIDULschedulingPDCCHCCEusage                    ProtocolIeID = 241
	ProtocolIeIDSFNOffset                                    ProtocolIeID = 242
	ProtocolIeIDQoSMonitoringDisabled                        ProtocolIeID = 243
	ProtocolIeIDExtendedUEIdentityIndexValue                 ProtocolIeID = 244
	ProtocolIeIDPagingeDRXInformation                        ProtocolIeID = 245
	ProtocolIeIDCHOMRDCEarlyDataForwarding                   ProtocolIeID = 246
	ProtocolIeIDSCGIndicator                                 ProtocolIeID = 247
	ProtocolIeIDUESpecificDRX                                ProtocolIeID = 248
	ProtocolIeIDPDUSessionExpectedUEActivityBehaviour        ProtocolIeID = 249
	ProtocolIeIDQoSMappingInformation                        ProtocolIeID = 250
	ProtocolIeIDAdditionLocationInformation                  ProtocolIeID = 251
	ProtocolIeIDdataForwardingInfoFromTargetEUTRANnode       ProtocolIeID = 252
	ProtocolIeIDDirectForwardingPathAvailability             ProtocolIeID = 253
	ProtocolIeIDSourceNGRANnodeID                            ProtocolIeID = 254
	ProtocolIeIDSourceDLForwardingIPAddress                  ProtocolIeID = 255
	ProtocolIeIDSourceNodeDLForwardingIPAddress              ProtocolIeID = 256
	ProtocolIeIDExtendedReportIntervalMDT                    ProtocolIeID = 257
	ProtocolIeIDSecurityIndication                           ProtocolIeID = 258
	ProtocolIeIDRRCConnReestabIndicator                      ProtocolIeID = 259
	ProtocolIeIDTargetNodeID                                 ProtocolIeID = 260
)