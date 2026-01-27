package email

// https://www.iana.org/assignments/imap-jmap-keywords/imap-jmap-keywords.xhtml

type Keyword = string

const (
	KeywordMDNSent       Keyword = "$MDNSent"
	KeywordForwarded     Keyword = "$Forwarded"
	KeywordSubmitPending Keyword = "$SubmitPending"
	KeywordSubmitted     Keyword = "$Submitted"
	KeywordJunk          Keyword = "$Junk"
	KeywordNotJunk       Keyword = "$NotJunk"
	KeywordPhishing      Keyword = "$Phishing"
	KeywordImportant     Keyword = "$Important"
	KeywordDraft         Keyword = "$draft"
	KeywordSeen          Keyword = "$seen"
	KeywordFlagged       Keyword = "$flagged"
	KeywordAnswered      Keyword = "$answered"
)
