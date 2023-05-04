<script setup lang="ts">
import { Message } from "src/constants/message";
import {
  MessageInterface,
  Message as MessagePayload,
  Request,
  HistoryRequest,
  HistoryPayload,
  FetchedHistoryTopic
} from "src/constants";
import { ref, Ref, watch, reactive, onMounted, computed } from "vue";
import { QScrollArea, useQuasar, Notify } from "quasar";
import {
  greetings,
  confusedResponse,
  helpfulResponse,
} from "src/constants/history";
import { useUtility } from "src/composables/useUtility";
import { useMessages } from "src/composables/useMessages";
import { botAvatar, userAvatar } from "src/constants/avatar";
import { api } from "src/boot/axios";
import { AxiosError } from "axios";

const $q = useQuasar();
const BREAKPOINT = 1024;
const isSmallScreen = computed(() => $q.screen.width < BREAKPOINT);

const splitter = computed(() => (isSmallScreen.value ? 0 : 30));
const drawer = ref(false);
const scrollArea = ref<QScrollArea | null>(null);
const scrollToBottom = () => {
  const scrollTarget = scrollArea.value?.getScrollTarget();
  const scrollPosition = scrollArea.value?.getScrollPosition();
  if (scrollTarget && scrollPosition) {
    scrollArea.value?.setScrollPosition(
      "vertical",
      scrollTarget.scrollHeight,
      1000
    );
  }
};

// This must be updated to the newest one (history length + 1)
const currentConversationID = ref(0);

// Perform fetching data to fill this array
const chatHistories: HistoryPayload = reactive({
  historyCollection: [],
});

const messages: MessageInterface[] = reactive([]);

// Fetch all history topic & id only
const fetchAllTopic = async () => {
  $q.loading.show({
    message: "Fetching important resouces. Hang on...",
  });
  const response = await api.get(
    "https://iridescent-jalebi-788066.netlify.app/.netlify/functions/endpoint/history-topic"
  );
  // console.log(response.data);
  if (response.data.historyPayload.HistoryCollection !== null) {
    const fetchedHistoryTopic: FetchedHistoryTopic[] = response.data.historyPayload.HistoryCollection;
    fetchedHistoryTopic.forEach((ele, index) => {
      const historyTopic: HistoryRequest = {
        historyId: ele.HistoryID,
        historyTopic: ele.HistoryTopic,
      }
      chatHistories.historyCollection.push(historyTopic);
    });
  } else{
    chatHistories.historyCollection = [];
  }
  // console.log(chatHistories.historyCollection);
  $q.loading.hide();
};

const fetchHistory = async () => {
  $q.loading.show({
    message: "Loading chat histories. Hang on...",
  });
  // console.log(currentConversationID.value);
  const response = await api.post(
    "https://iridescent-jalebi-788066.netlify.app/.netlify/functions/endpoint/history",
    { id: currentConversationID.value }
  );
  console.log(response.data);
  const conversation: [] = response.data.history.conversation;
  messages.splice(0, messages.length);
  conversation.forEach((message: MessagePayload, index) => {
    const tempMessage = new Message(
      message.id,
      message.sent,
      message.text,
      message.sentTime,
      message.historyId
    )
    tempMessage.setResponse(message.response, 200);
    tempMessage.setResponseStatus(true);
    console.log(tempMessage);
    messages.push(tempMessage);
  }) 
  messages.forEach((message, index) => {
    message.text = message.text.replace(/\n/g, "<br>");
  });
  console.log(messages);
  $q.loading.hide();
};

const userInput: Ref<string> = ref("");
const botResponse: Ref<string> = ref("");
const botFullResponse: Ref<string> = ref("");
const isResponding: Ref<boolean> = ref(false);
const showGPTinfo: Ref<boolean> = ref(true);

const { animateMessage, random, generateTimestamp } = useUtility({
  startNum: 0,
  endNum: 5,
  botMessage: botResponse,
  duration: 20,
});

const botGreetings = ref(greetings[random()]);
const method = ref("KMP");

const newChat = () => {
  showGPTinfo.value = true;
  botGreetings.value = greetings[random()];
  currentConversationID.value = chatHistories.historyCollection.length + 1;
  messages.splice(0, messages.length);
};

const switchConversation = async () => {
  // console.log(currentConversationID.value);
  // replace the current array to chosen history conversation
  await fetchHistory();
  scrollToBottom();
};

const sendMessage = async () => {
  if (!isResponding.value) {
    if (showGPTinfo.value) {
      showGPTinfo.value = false;
    }
    const filteredStr = userInput.value.replace(/\n/g, "");
    const { generateMessageId } = useMessages({ chatHistories });
    isResponding.value = true;
    scrollToBottom();
    let userMessage: Message;
    if (messages.length === 0) {
      const availId = generateMessageId();
      userMessage = new Message(
        messages.length + 1,
        true,
        filteredStr,
        new Date().toLocaleTimeString(),
        availId
      );
      currentConversationID.value = availId;
      userMessage.setHistoryTimestamp(generateTimestamp());
    } else {
      userMessage = new Message(
        messages.length + 1,
        true,
        filteredStr,
        new Date().toLocaleTimeString(),
        currentConversationID.value
      );
    }
    messages.push(userMessage);
    // const currentTopic = filteredStr;

    // Send request to backend
    const request: Request = {
      id: userMessage.getId(),
      text: userMessage.getText(),
      response: "",
      sentTime: userMessage.getSentTime(),
      historyId: userMessage.getHistoryId(),
      historyTimestamp: userMessage.getHistoryTimestamp(),

      method: method.value as "KMP" | "BoyerMoore" | "GPT",
    };

    // Clear input
    userInput.value = "";
    if (method.value !== "GPT") {
      try {
        // Fetch response from backend
        const response = await api.post(
          "https://iridescent-jalebi-788066.netlify.app/.netlify/functions/endpoint/getmessage",
          request
        );
        // Unload response from api
        if (response.data.message === "200") {
          botFullResponse.value = response.data.botResponse.response + "\n\n";
          botFullResponse.value += helpfulResponse[random()];
        } else {
          // const confusedText = confusedResponse[random()];
          // const choiceArr = response.data.botResponse.response.split("<-|->");
          // botFullResponse.value = confusedText;
          // choiceArr.forEach((choice: string, index: number) => {
          //   if (choice !== "") {
          //     botFullResponse.value += index + 1 + ". " + choice + "\n";
          //   }
          // });
          botFullResponse.value = response.data.botResponse.response;
          botFullResponse.value +=
            "Please rewrite the question if you desire to choose from the above!";
        }
        botFullResponse.value = botFullResponse.value.replace(/\n/g, "<br>");
      } catch (error) {
        const axiosError = error as AxiosError;
        if (axiosError.response) {
          switch (axiosError.response.status) {
            case 502:
              let i = 5;
              const intervalId = setInterval(() => {
                Notify.create({
                  message: `Something wrong with our end! Please come back later!`,
                });
                i--;
                if (i < 1) {
                  clearInterval(intervalId);
                }
              }, 1000);
              break;
          }
        }
      }
    } else {
      const response = await api.post(
        "https://tubes3stima.pythonanywhere.com/completion",
        { request: request.text },
        {
          headers: { "Content-Type": "application/json" },
        }
      );
      if (response !== undefined) {
        botFullResponse.value = response.data + "\n\n";
        botFullResponse.value += helpfulResponse[random()];
      }
    }

    messages[messages.length - 1].setResponse(botFullResponse.value, 200);
    await animateMessage(botFullResponse.value);
    messages[messages.length - 1].setResponseStatus(true);
    botResponse.value = "";
    isResponding.value = false;

    // if (messages.length === 1) {
    //   const currentMessages = messages.slice();
    //   const newHistory: History = {
    //     historyId: generateMessageId(),
    //     topic: currentTopic,
    //     conversation: currentMessages,
    //   };
    //   chatHistories.messageHistory.push(newHistory);
    // } else {
    //   updateHistory(currentConversationID.value, userMessage);
    // }
  }
};

watch([messages, currentConversationID], () => {
  scrollToBottom();
});

onMounted(async () => {
  scrollToBottom();
  await fetchAllTopic();
  const { generateMessageId } = useMessages({ chatHistories });
  currentConversationID.value = generateMessageId();
});
</script>
<template>
  <q-page class="tw-h-screen bg-primary">
    <q-drawer v-model="drawer" overlay bordered>
      <div class="q-pa-md tw-h-screen bg-grey-2">
        <div class="text-h4 q-mb-md text-accent tw-flex tw-justify-center">
          Chat History
        </div>
        <q-scroll-area
          style="height: 70%"
          class="tw-px-1"
          :vertical-thumb-style="{ width: '5px' }"
        >
          <q-btn
            dark
            outline
            label="New Chat"
            icon="add_circle"
            class="tw-w-full tw-h-12 text-accent tw-rounded-2xl"
            @click="newChat"
            :disable="isResponding"
          />

          <q-tabs
            v-model="currentConversationID"
            vertical
            class="tw-py-4 text-body text-accent tw-w-full"
            active-bg-color="secondary"
            active-color="white"
          >
            <template
              v-for="history in chatHistories.historyCollection"
              :key="history"
            >
              <q-tab
                :name="history.historyId"
                @click="switchConversation"
                style="border-bottom: 2px solid #46b1c9"
                content-class="tw-w-full"
                :disable="isResponding"
              >
                <div class="tw-flex tw-items-center">
                  <q-icon name="chat" class="tw-pr-2" />
                  <template v-if="history.historyTopic.length > 15">
                    <span>{{ history.historyTopic.substring(0, 15) }}...</span>
                  </template>
                  <template v-else>
                    <span>{{ history.historyTopic.substring(0, 15) }}</span>
                  </template>
                </div>
              </q-tab>
            </template>
          </q-tabs>
        </q-scroll-area>
        <div class="tw-mt-1 text-accent">
          <h3>Algorithm:</h3>
          <q-radio
            v-model="method"
            checked-icon="task_alt"
            unchecked-icon="panorama_fish_eye"
            color="accent"
            val="KMP"
            label="KMP"
            :disable="isResponding"
            class="-tw-mt-10"
          />
          <q-radio
            v-model="method"
            checked-icon="task_alt"
            unchecked-icon="panorama_fish_eye"
            color="accent"
            val="BoyerMoore"
            label="Boyer Moore"
            :disable="isResponding"
            class="-tw-mt-10"
          />
          <q-radio
            v-model="method"
            checked-icon="task_alt"
            unchecked-icon="panorama_fish_eye"
            color="accent"
            val="GPT"
            label="GPT"
            :disable="isResponding"
            class="-tw-mt-10"
          />
        </div>
        <span class="text-sm-caption text-black tw-absolute tw-top-1 tw-left-2">
          Copyright by 666
        </span>
      </div>
    </q-drawer>
    <q-splitter v-model="splitter" :separator-style="{ display: 'none' }">
      <template v-slot:before>
        <div
          class="q-pa-md tw-h-screen tw-max-w-lg tw-pr-1"
          style="border-right: solid 2px #f46197"
        >
          <div class="text-h4 q-mb-md text-accent tw-mt-2">Chat History</div>
          <q-scroll-area
            style="height: 70%"
            class="tw-px-2"
            :vertical-thumb-style="{ width: '5px' }"
          >
            <q-btn
              dark
              outline
              label="New Chat"
              icon="add_circle"
              class="tw-w-full tw-h-12 text-accent tw-rounded-2xl"
              @click="newChat"
              :disable="isResponding"
            />

            <q-tabs
              v-model="currentConversationID"
              vertical
              class="tw-py-4 text-body text-accent"
              active-bg-color="secondary"
              active-color="white"
            >
              <template
                v-for="history in chatHistories.historyCollection"
                :key="history"
              >
                <q-tab
                  :name="history.historyId"
                  @click="switchConversation"
                  style="border-bottom: 2px solid #46b1c9"
                  content-class="tw-w-full"
                  :disable="isResponding"
                >
                  <div class="tw-flex tw-items-center">
                    <q-icon name="chat" class="tw-pr-2" />
                    <template v-if="history.historyTopic.length > 20">
                      <span>{{ history.historyTopic.substring(0, 20) }}...</span>
                    </template>
                    <template v-else>
                      <span>{{ history.historyTopic.substring(0, 20) }}</span>
                    </template>
                  </div>
                </q-tab>
              </template>
            </q-tabs>
          </q-scroll-area>
          <div class="tw-mt-10 text-accent">
            <h3>Algorithm:</h3>
            <q-radio
              dark
              v-model="method"
              checked-icon="task_alt"
              unchecked-icon="panorama_fish_eye"
              color="accent"
              val="KMP"
              label="KMP"
              :disable="isResponding"
            />
            <q-radio
              dark
              v-model="method"
              checked-icon="task_alt"
              unchecked-icon="panorama_fish_eye"
              color="accent"
              val="BoyerMoore"
              label="Boyer Moore"
              :disable="isResponding"
            />
            <q-radio
              dark
              v-model="method"
              checked-icon="task_alt"
              unchecked-icon="panorama_fish_eye"
              color="accent"
              val="GPT"
              label="GPT"
              :disable="isResponding"
            />
          </div>
          <span class="text-caption text-white tw-absolute tw-top-0 tw-left-4">
            Copyright by 666
          </span>
        </div>
      </template>

      <template v-slot:after>
        <div
          style="width: 100%"
          class="tw-h-screen tw-w-flex tw-flex-col tw-justify-end"
        >
          <q-btn
            flat
            v-if="isSmallScreen"
            icon="menu"
            size="md"
            class="tw-absolute bg-info tw-m-4 tw-z-10"
            @click="() => (drawer = !drawer)"
          />
          <q-scroll-area
            style="height: 90%"
            :delay="1000"
            class="tw-px-5 chat-area"
            ref="scrollArea"
          >
            <template v-if="method === 'GPT' && showGPTinfo">
              <h6 class="tw-absolute text-warning tw-text-center tw-w-full tw-top-52">
                Welcome to GPT Mode!
                <br>
                Please note that GPT Mode will not be saved to our database!
              </h6>
            </template>
            <div class="tw-w-full">
              <q-chat-message
                name="BOT"
                :avatar="botAvatar"
                bg-color="grey-1"
                text-color="white"
              >
                <div class="text-lg-body">{{ botGreetings }}</div>
              </q-chat-message>
              <div v-for="message in messages" :key="message.id">
                <q-chat-message
                  :sent="true"
                  text-color="white"
                  bg-color="green-1"
                  :avatar="userAvatar"
                  :stamp="message.getSentTime()"
                >
                  <div class="text-lg-body">{{ message.getText() }}</div>
                </q-chat-message>
                <q-chat-message
                  name="BOT"
                  :avatar="botAvatar"
                  bg-color="grey-1"
                  text-color="white"
                >
                  <template v-if="message.getResponseCode() === 0">
                    <q-spinner-dots size="2rem" />
                  </template>
                  <template v-else>
                    <div class="text-lg-body">
                      <template
                        v-if="
                          botResponse.length !== botFullResponse.length &&
                          !message.getResponseStatus()
                        "
                      >
                        <div v-html="botResponse"></div>
                      </template>
                      <template v-else>
                        <div v-html="message.getResponseMsg()"></div>
                      </template>
                    </div>
                  </template>
                </q-chat-message>
              </div>
            </div>
          </q-scroll-area>
          <div class="row items-center tw-gap-x-2 tw-px-2 tw-mt-4">
            <q-input
              dark
              v-model="userInput"
              label="Type your message here"
              class="tw-grow"
              @keyup.enter="sendMessage"
              autogrow
            />
            <div>
              <q-btn
                flat
                icon="send"
                @click="sendMessage"
                size="md"
                round
                class="text-green-1"
                :disable="
                  isResponding ||
                  userInput.trim().length === 0 ||
                  userInput === undefined
                "
              />
            </div>
          </div>
        </div>
      </template>
    </q-splitter>
  </q-page>
</template>

<style lang="scss" scoped>
.chat-area {
  :deep(.q-scrollarea__content) {
    display: flex;
    width: 100% !important;
    align-items: flex-end !important;
    position: relative;
  }
}

:deep(.q-message-name) {
  color: white;
}

:deep(.q-btn__content) {
  font-size: 16px;
  text-transform: none !important;
}

:deep(.q-field__control-container) {
  padding-inline: 10px;
}
:deep(.q-field__label) {
  margin-left: 10px !important;
  color: white;
}
:deep(.q-field__native) {
  font-size: 20px;
}
:deep(.q-tab__indicator) {
  display: none !important;
}
.q-message {
  animation: scale-in 0.5s;
}

@keyframes scale-in {
  from {
    transform: scale(0);
  }
  to {
    transform: scale(1);
  }
}
</style>
