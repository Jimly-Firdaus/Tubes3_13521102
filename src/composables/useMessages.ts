import { Message } from "src/constants/message";
import { MessageHistory } from "src/constants";
// import { allHistory } from "src/constants/history";

export const useMessages = ({
  chatHistories,
}: {
  chatHistories?: MessageHistory;
} = {}) => {
  const generateMessageId = () => {
    if (chatHistories) {
      chatHistories.messageHistory.forEach((history, index) => {
        if (history.historyId !== index + 1) {
          return index + 1;
        }
      });
      return chatHistories.messageHistory.length + 1;
    } else {
      return -1;
    }
  };

  const updateHistory = (historyId: number, message: Message) => {
    if (chatHistories) {
      chatHistories.messageHistory.forEach((history, index) => {
        if (history.historyId === historyId) {
          history.conversation.push(message);
          return;
        }
      });
    }
  };

  return {
    generateMessageId,
    updateHistory,
  };
};
