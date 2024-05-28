package utils

import (
	"log"

	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/response"
	"github.com/lib/pq"
)

func HandleDBError(err error) string {
	if pqErr, ok := err.(*pq.Error); ok {
		log.Printf("PostgreSQL Error Code: %s", pqErr.Code)

		switch pqErr.Code {
		case response.SuccessfulCompletion:
			return "Successful completion"
		case response.Warning:
			return "Warning"
		case response.DynamicResultSetsReturned:
			return "Dynamic result sets returned"
		case response.ImplicitZeroBitPadding:
			return "Implicit zero bit padding"
		case response.NullValueEliminatedInSetFunction:
			return "Null value eliminated in set function"
		case response.PrivilegeNotGranted:
			return "Privilege not granted"
		case response.PrivilegeNotRevoked:
			return "Privilege not revoked"
		case response.StringDataRightTruncation:
			return "String data right truncation"
		case response.DeprecatedFeature:
			return "Deprecated feature"
		case response.NoData:
			return "No data"
		case response.NoAdditionalDynamicResultSetsReturned:
			return "No additional dynamic result sets returned"
		case response.SQLStatementNotYetComplete:
			return "SQL statement not yet complete"
		case response.ConnectionException:
			return "Connection exception"
		case response.ConnectionDoesNotExist:
			return "Connection does not exist"
		case response.ConnectionFailure:
			return "Connection failure"
		case response.SQLClientUnableToEstablishSQLConnection:
			return "SQL client unable to establish SQL connection"
		case response.SQLServerRejectedEstablishmentOfSQLConnection:
			return "SQL server rejected establishment of SQL connection"
		case response.TransactionResolutionUnknown:
			return "Transaction resolution unknown"
		case response.ProtocolViolation:
			return "Protocol violation"
		case response.TriggeredActionException:
			return "Triggered action exception"
		case response.FeatureNotSupported:
			return "Feature not supported"
		case response.InvalidTransactionInitiation:
			return "Invalid transaction initiation"
		case response.LocatorException:
			return "Locator exception"
		case response.InvalidLocatorSpecification:
			return "Invalid locator specification"
		case response.InvalidGrantor:
			return "Invalid grantor"
		case response.InvalidGrantOperation:
			return "Invalid grant operation"
		case response.InvalidRoleSpecification:
			return "Invalid role specification"
		case response.DiagnosticsException:
			return "Diagnostics exception"
		case response.StackedDiagnosticsAccessedWithoutActiveHandler:
			return "Stacked diagnostics accessed without active handler"
		case response.CaseNotFound:
			return "Case not found"
		case response.CardinalityViolation:
			return "Cardinality violation"
		case response.DataException:
			return "Data exception"
		case response.ArraySubscriptError:
			return "Array subscript error"
		case response.CharacterNotInRepertoire:
			return "Character not in repertoire"
		case response.DivisionByZero:
			return "Division by zero"
		case response.ErrorInAssignment:
			return "Error in assignment"
		case response.EscapeCharacterConflict:
			return "Escape character conflict"
		case response.IndicatorOverflow:
			return "Indicator overflow"
		case response.IntervalFieldOverflow:
			return "Interval field overflow"
		case response.InvalidArgumentForLogarithm:
			return "Invalid argument for logarithm"
		case response.InvalidArgumentForNtileFunction:
			return "Invalid argument for ntile function"
		case response.InvalidArgumentForNthValueFunction:
			return "Invalid argument for nth value function"
		case response.InvalidArgumentForPowerFunction:
			return "Invalid argument for power function"
		case response.InvalidArgumentForWidthBucketFunction:
			return "Invalid argument for width bucket function"
		case response.InvalidCharacterValueForCast:
			return "Invalid character value for cast"
		case response.InvalidDatetimeFormat:
			return "Invalid datetime format"
		case response.InvalidEscapeCharacter:
			return "Invalid escape character"
		case response.InvalidEscapeOctet:
			return "Invalid escape octet"
		case response.InvalidEscapeSequence:
			return "Invalid escape sequence"
		case response.NonstandardUseOfEscapeCharacter:
			return "Nonstandard use of escape character"
		case response.InvalidIndicatorParameterValue:
			return "Invalid indicator parameter value"
		case response.InvalidParameterValue:
			return "Invalid parameter value"
		case response.InvalidRegularExpression:
			return "Invalid regular expression"
		case response.InvalidRowCountInLimitClause:
			return "Invalid row count in limit clause"
		case response.InvalidRowCountInResultOffsetClause:
			return "Invalid row count in result offset clause"
		case response.InvalidTimeZoneDisplacementValue:
			return "Invalid time zone displacement value"
		case response.InvalidUseOfEscapeCharacter:
			return "Invalid use of escape character"
		case response.MostSpecificTypeMismatch:
			return "Most specific type mismatch"
		case response.NullValueNotAllowed:
			return "Null value not allowed"
		case response.NullValueNoIndicatorParameter:
			return "Null value no indicator parameter"
		case response.NumericValueOutOfRange:
			return "Numeric value out of range"
		case response.SequenceGeneratorLimitExceeded:
			return "Sequence generator limit exceeded"
		case response.StringDataLengthMismatch:
			return "String data length mismatch"
		case response.SubstringError:
			return "Substring error"
		case response.TrimError:
			return "Trim error"
		case response.UnterminatedCString:
			return "Unterminated C string"
		case response.ZeroLengthCharacterString:
			return "Zero length character string"
		case response.FloatingPointException:
			return "Floating point exception"
		case response.InvalidTextRepresentation:
			return "Invalid text representation"
		case response.InvalidBinaryRepresentation:
			return "Invalid binary representation"
		case response.BadCopyFileFormat:
			return "Bad copy file format"
		case response.UntranslatableCharacter:
			return "Untranslatable character"
		case response.NotAnXMLDocument:
			return "Not an XML document"
		case response.InvalidXMLDocument:
			return "Invalid XML document"
		case response.InvalidXMLContent:
			return "Invalid XML content"
		case response.InvalidXMLComment:
			return "Invalid XML comment"
		case response.InvalidXMLProcessingInstruction:
			return "Invalid XML processing instruction"
		case response.IntegrityConstraintViolation:
			return "Integrity constraint violation"
		case response.RestrictViolation:
			return "Restrict violation"
		case response.NotNullViolation:
			return "Not null violation"
		case response.ForeignKeyViolation:
			return "Foreign key violation"
		case response.UniqueViolation:
			return "Unique violation"
		case response.CheckViolation:
			return "Check violation"
		case response.ExclusionViolation:
			return "Exclusion violation"
		case response.InvalidCursorState:
			return "Invalid cursor state"
		case response.InvalidTransactionState:
			return "Invalid transaction state"
		case response.ActiveSQLTransaction:
			return "Active SQL transaction"
		case response.BranchTransactionAlreadyActive:
			return "Branch transaction already active"
		case response.HeldCursorRequiresSameIsolationLevel:
			return "Held cursor requires same isolation level"
		case response.InappropriateAccessModeForBranchTransaction:
			return "Inappropriate access mode for branch transaction"
		case response.InappropriateIsolationLevelForBranchTransaction:
			return "Inappropriate isolation level for branch transaction"
		case response.NoActiveSQLTransactionForBranchTransaction:
			return "No active SQL transaction for branch transaction"
		case response.ReadOnlySQLTransaction:
			return "Read-only SQL transaction"
		case response.SchemaAndDataStatementMixingNotSupported:
			return "Schema and data statement mixing not supported"
		case response.NoActiveSQLTransaction:
			return "No active SQL transaction"
		case response.InFailedSQLTransaction:
			return "In failed SQL transaction"
		case response.InvalidSQLStatementName:
			return "Invalid SQL statement name"
		case response.TriggeredDataChangeViolation:
			return "Triggered data change violation"
		case response.InvalidAuthorizationSpecification:
			return "Invalid authorization specification"
		case response.InvalidPassword:
			return "Invalid password"
		case response.DependentPrivilegeDescriptorsStillExist:
			return "Dependent privilege descriptors still exist"
		case response.DependentObjectsStillExist:
			return "Dependent objects still exist"
		case response.InvalidTransactionTermination:
			return "Invalid transaction termination"
		case response.SQLRoutineException:
			return "SQL routine exception"
		case response.FunctionExecutedNoReturnStatement:
			return "Function executed no return statement"
		case response.ModifyingSQLDataNotPermitted:
			return "Modifying SQL data not permitted"
		case response.ProhibitedSQLStatementAttempted:
			return "Prohibited SQL statement attempted"
		case response.ReadingSQLDataNotPermitted:
			return "Reading SQL data not permitted"
		case response.InvalidCursorName:
			return "Invalid cursor name"
		case response.ExternalRoutineException:
			return "External routine exception"
		case response.ContainingSQLNotPermitted:
			return "Containing SQL not permitted"
		case response.ExternalRoutineInvocationException:
			return "External routine invocation exception"
		case response.InvalidSQLStateReturned:
			return "Invalid SQL state returned"
		case response.TriggerProtocolViolated:
			return "Trigger protocol violated"
		case response.SRFProtocolViolated:
			return "SRF protocol violated"
		case response.SavepointException:
			return "Savepoint exception"
		case response.InvalidSavepointSpecification:
			return "Invalid savepoint specification"
		case response.InvalidCatalogName:
			return "Invalid catalog name"
		case response.InvalidSchemaName:
			return "Invalid schema name"
		case response.TransactionRollback:
			return "Transaction rollback"
		case response.TransactionIntegrityConstraintViolation:
			return "Transaction integrity constraint violation"
		case response.SerializationFailure:
			return "Serialization failure"
		case response.StatementCompletionUnknown:
			return "Statement completion unknown"
		case response.DeadlockDetected:
			return "Deadlock detected"
		case response.SyntaxErrorOrAccessRuleViolation:
			return "Syntax error or access rule violation"
		case response.SyntaxError:
			return "Syntax error"
		case response.InsufficientPrivilege:
			return "Insufficient privilege"
		case response.CannotCoerce:
			return "Cannot coerce"
		case response.GroupingError:
			return "Grouping error"
		case response.WindowingError:
			return "Windowing error"
		case response.InvalidRecursion:
			return "Invalid recursion"
		case response.InvalidForeignKey:
			return "Invalid foreign key"
		case response.InvalidName:
			return "Invalid name"
		case response.NameTooLong:
			return "Name too long"
		default:
			return "No PostgreSQL error"
		}
	}
	return ""
}
