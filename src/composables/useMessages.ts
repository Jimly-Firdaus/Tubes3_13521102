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
      const historyIds = chatHistories.messageHistory.map(
        (history) => history.historyId
      );
      historyIds.sort((a, b) => a - b);
      let availableId = 1;
      for (let i = 0; i < historyIds.length; i++) {
        if (historyIds[i] !== availableId) {
          break;
        }
        availableId++;
      }
      return availableId;
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
