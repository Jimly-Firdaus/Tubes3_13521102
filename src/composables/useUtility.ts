import { Ref, WritableComputedRef } from "vue";
export const useUtility = ({
  startNum,
  endNum,
  botMessage,
  duration,
}: {
  startNum?: number;
  endNum?: number;
  botMessage?: Ref<string> | WritableComputedRef<string>;
  duration?: number;
} = {}) => {
  /**
   * Wait for duration in millis
   */
  const wait = () => new Promise((resolve) => setTimeout(resolve, duration));

  /**
   * Animate bot message per word to full message
   * */
  const animateMessage = async (botFullMessage: string): Promise<void> => {
    let i = 0;
    if (botMessage) {
      while (i <= botFullMessage.length) {
        botMessage.value +=
          botFullMessage.charAt(i) + botFullMessage.charAt(i + 1);
        i = i + 2;
        await wait();
      }
    }
  };

  /**
   * Give random number form startNum to endNum if given
   * otherwise return -1
   * */
  const random = (): number => {
    if (startNum && endNum) {
      return Math.floor(Math.random() * endNum) + startNum;
    }
    return -1;
  };

  /**
   * Generate timestamp for current time
   * @returns formatted date (YYYY-MM-DD HH:MM:SS)
   * */
  const generateTimestamp = (): string => {
    const date = new Date();
    const year = date.getFullYear();
    const month = date.getMonth() + 1;
    const day = date.getDate();
    const hours = date.getHours();
    const minutes = date.getMinutes();
    const seconds = date.getSeconds();

    const formattedDate = `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`;
    return formattedDate;
  };

  return {
    animateMessage,
    random,
    generateTimestamp,
  };
};
