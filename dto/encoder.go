package dto

import (
	"net/url"

	"github.com/sonh/qs"
)

type hasQSTag interface {
	UserListRequest |
		UserGetRequest |
		UserListGroupChannelRequest |
		OpenChannelListRequest |
		OpenChannelListOperatorsRequest |
		OpenChannelListParticipantsRequest |
		OpenChannelUnregisterOperatorRequest |
		GroupChannelListRequest |
		GroupChannelListOperatorsRequest |
		GroupChannelListMembersRequest |
		GroupChannelUnregisterOperatorRequest |
		CustomChannelListSettingsRequest |
		CustomChannelListOperatorsRequest |
		CustomChannelUnregisterOperatorRequest |
		MessageListRequest |
		MessageListThreadedRepliesRequest |
		MessageGetRequest |
		MessageDeleteMetaRequest |
		MessageViewReplyRequest |
		MessageViewThreadInfoRequest |
		EventSubscriptionListRequest |
		DataExportRequest |
		ListReportRequest |
		ListReportUserRequest |
		ListReportChannelRequest |
		ListReportMessageRequest |
		ListReportModerationMessageRequest
}

// encodeQS takes baseURL and struct with `qs` tag
// encodes struct with qs tag to URL string
// will return the encoded url string and error if any
func encodeQS[T hasQSTag](baseURL string, input T) (string, error) {
	_, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}

	encoder := qs.NewEncoder()
	values, err := encoder.Values(input)
	if err != nil {
		return "", err
	}

	return baseURL + values.Encode(), nil
}
