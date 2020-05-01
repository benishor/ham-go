// Code generated by "stringer -type=MessageCode"; DO NOT EDIT.

package wsjtx

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[MessageHeartbeat-0]
	_ = x[MessageStatus-1]
	_ = x[MessageDecode-2]
	_ = x[MessageClear-3]
	_ = x[MessageReply-4]
	_ = x[MessageQSOLogged-5]
	_ = x[MessageClose-6]
	_ = x[MessageReplay-7]
	_ = x[MessageHaltTX-8]
	_ = x[MessageFreeText-9]
	_ = x[MessageWSPRDecode-10]
	_ = x[MessageLocation-11]
	_ = x[MessageLoggedADIF-12]
}

const _MessageCode_name = "MessageHeartbeatMessageStatusMessageDecodeMessageClearMessageReplyMessageQSOLoggedMessageCloseMessageReplayMessageHaltTXMessageFreeTextMessageWSPRDecodeMessageLocationMessageLoggedADIF"

var _MessageCode_index = [...]uint8{0, 16, 29, 42, 54, 66, 82, 94, 107, 120, 135, 152, 167, 184}

func (i MessageCode) String() string {
	if i >= MessageCode(len(_MessageCode_index)-1) {
		return "MessageCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _MessageCode_name[_MessageCode_index[i]:_MessageCode_index[i+1]]
}
