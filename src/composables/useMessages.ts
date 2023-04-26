import { MessageHistory } from "src/constants";
// import { allHistory } from "src/constants/history";

export const useMessages = ({
  allHistory
}: {
  allHistory?: MessageHistory;
} = {}) => {
    const generateMessageId = () => {
        if(allHistory) {
            allHistory.messageHistory.forEach((history, index) => {
                if (history.historyId !== index + 1) {
                    return index + 1;
                }
            })
            return allHistory.messageHistory.length + 1;
        }
        else {
            return -1;
        }
    } 

  return {
    generateMessageId
  };
};