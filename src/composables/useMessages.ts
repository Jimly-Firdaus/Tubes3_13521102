import { Message } from "src/constants/message";
import { MessageHistory, HistoryPayload } from "src/constants";
// import { allHistory } from "src/constants/history";

export const useMessages = ({
  chatHistories,
}: {
  chatHistories?: HistoryPayload;
} = {}) => {
  const generateMessageId = () => {
    if (chatHistories) {
      const historyIds = chatHistories.historyCollection.map(
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

  return {
    generateMessageId,
  };
};
