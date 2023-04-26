import { MessageHistory, History, Message } from "src/constants/index"
import { Message as MessageConstructor } from "./message";

export const chatTopic: string[] = [
  "Topic - 1",
  "Topic - 2",
  "Topic - 3",
  "Topic - 4",
  "Topic - 5",
  "Topic - 6",
  "Topic - 7",
  "Topic - 8",
  "Topic - 9",
];

export const dummyResponse: string[] = [
  "Lorem ipsum dolor sit, amet consectetur adipisicing elit. Quis praesentium cumque magnam odio iure quidem, quod illum numquam possimus obcaecati commodi minima assumenda consectetur culpa fuga nulla ullam. In, libero.",
  "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Egestas sed tempus urna et pharetra pharetra.",
  "Non pulvinar neque laoreet suspendisse interdum. Sed elementum tempus egestas sed. Cum sociis natoque penatibus et magnis dis parturient montes nascetur.",
  "Fermentum leo vel orci porta non pulvinar. Nisl nunc mi ipsum faucibus vitae aliquet nec ullamcorper sit. Porttitor lacus luctus accumsan tortor. Semper risus in hendrerit gravida rutrum.",
  "At elementum eu facilisis sed odio morbi quis commodo odio. Ultricies tristique nulla aliquet enim tortor. Faucibus et molestie ac feugiat.",
  "Nisl pretium fusce id velit ut tortor pretium viverra suspendisse. Semper risus in hendrerit gravida rutrum quisque. Est lorem ipsum dolor sit amet consectetur adipiscing elit pellentesque.",
  "Massa enim nec dui nunc mattis enim ut tellus elementum. Donec ultrices tincidunt arcu non sodales neque sodales ut.",
  "Eget nunc scelerisque viverra mauris in aliquam sem fringilla. Tellus elementum sagittis vitae et leo duis ut diam.",
];

export const message1: MessageConstructor[] = [];
export const message2: MessageConstructor[] = [];

for (let i = 0; i < 14; i++) {
  const message = new MessageConstructor(
    i,
    true,
    `This message is from message1 array ${i + 1}`,
    new Date().toLocaleTimeString(),
    1
  )
  message.setResponse(`Response for message1 array ${i + 1}`, 200);
  message1.push(message);
}

for (let i = 0; i < 14; i++) {
  const message = new MessageConstructor(
    i,
    true,
    `This message is from message2 array ${i + 1}`,
    new Date().toLocaleTimeString(),
    2
  )
  message.setResponse(`Response for message2 array ${i + 1}`, 200);
  message2.push(message);
}

export const history1: History = {
  historyId: 1,
  topic: "Message - 1 Array",
  conversation: message1
}

export const history2: History = {
  historyId: 2,
  topic: "Message - 2 Array",
  conversation: message2
}

export const allHistory: MessageHistory = {
  messageHistory: [history1, history2]
};

