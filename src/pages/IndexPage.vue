<script setup lang="ts">
import { Message } from "src/constants/message";
import {
  MessageInterface,
  History,
  Request,
  MessageHistory,
} from "src/constants";
import { ref, Ref, watch, reactive, onMounted, computed } from "vue";
import { QScrollArea, useQuasar } from "quasar";
import { dummyResponse } from "src/constants/history";
import { useUtility } from "src/composables/useUtility";
import { useMessages } from "src/composables/useMessages";
import { botAvatar, userAvatar } from "src/constants/avatar";
import { api } from "src/boot/axios";

const $q = useQuasar();
const BREAKPOINT = 1024;
const isSmallScreen = computed(() => $q.screen.width < BREAKPOINT);

const splitter = computed(() => (isSmallScreen.value ? 0 : 20));
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
const chatHistories: MessageHistory = reactive({
  messageHistory: [],
});
const messages: MessageInterface[] = reactive([]);

const fetchHistories = async () => {
  $q.loading.show({
    message: 'Fetching important resouces. Hang on...'
  });
  const response = await api.get("http://localhost:8080/history");
  const fetchedMessageHistory: [] = response.data.historyMessage.messageHistory;
  fetchedMessageHistory.forEach((ele: History, index) => {
    // console.log(ele);
    const arrayOfConversation: Array<MessageInterface> = [];
    ele.conversation.forEach((messageJSON, index) => {
      const message = new Message(
        messageJSON.id,
        true,
        messageJSON.text,
        messageJSON.sentTime,
        messageJSON.historyId
      )
      message.setResponse(messageJSON.response, 200);
      message.setResponseStatus(true);
      arrayOfConversation.push(message);
    })
    const history: History = {
      historyId: ele.historyId,
      topic: ele.topic,
      conversation: arrayOfConversation
    };
    chatHistories.messageHistory.push(history);
  })
  $q.loading.hide();
};

const userInput: Ref<string> = ref("");
const botResponse: Ref<string> = ref("");
const botFullResponse: Ref<string> = ref("");
const isResponding: Ref<boolean> = ref(false);

const { animateMessage, random, generateTimestamp } = useUtility({
  startNum: 1,
  endNum: 6,
  botMessage: botResponse,
  duration: 20,
});

const method = ref("KMP");

const newChat = () => {
  currentConversationID.value = chatHistories.messageHistory.length + 1;
  messages.splice(0, messages.length);
};

const switchConversation = () => {
  // replace the current array to chosen history conversation
  const chosenHistory =
    chatHistories.messageHistory[currentConversationID.value - 1];
  messages.splice(0, messages.length, ...chosenHistory.conversation);
  scrollToBottom();
};

const sendMessage = async () => {
  if (!isResponding.value) {
    const filteredStr = userInput.value.replace(/\n/g, "");
    const { generateMessageId, updateHistory } = useMessages({ chatHistories });
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
    const currentTopic = filteredStr;

    // Send request to backend
    const request: Request = {

      id: userMessage.getId(),
      text: userMessage.getText(),
      response: "",
      sentTime: userMessage.getSentTime(),
      historyId: userMessage.getHistoryId(),
      historyTimestamp: userMessage.getHistoryTimestamp(),


      method: method.value as "KMP" | "BoyerMoore",
    };
    // console.log(request);
    // Unload response from api
    const response = await api.post("http://localhost:8080/getmessage", request);
    console.log(response.data);

    // Clear input
    userInput.value = "";

    setTimeout(async () => {
      // Fetch response from backend
      botFullResponse.value = dummyResponse[random()];

      // Store the result
      messages[messages.length - 1].setResponse(botFullResponse.value, 200);

      // Show result
      await animateMessage(botFullResponse.value);
      messages[messages.length - 1].setResponseStatus(true);
      botResponse.value = "";
      isResponding.value = false;
    }, 1500);
    if (messages.length === 1) {
      const currentMessages = messages.slice();
      const newHistory: History = {
        historyId: generateMessageId(),
        topic: currentTopic,
        conversation: currentMessages,
      };
      chatHistories.messageHistory.push(newHistory);
    } else {
      updateHistory(currentConversationID.value, userMessage);
    }
  }
};

watch([messages, currentConversationID], () => {
  scrollToBottom();
});

onMounted(() => {
  scrollToBottom();
  fetchHistories();
});
</script>
<template>
  <q-page class="tw-h-screen bg-primary">
    <q-drawer v-model="drawer" overlay bordered>
      <div class="q-pa-md tw-h-screen bg-grey-2">
        <div class="text-h4 q-mb-md text-accent">Chat History</div>
        <q-scroll-area
          style="height: 70%"
          class="tw-pr-2"
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
              v-for="history in chatHistories.messageHistory"
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
                  <span>{{ history.topic }}</span>
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
        </div>
        <span class="text-sm-caption text-black tw-absolute tw-bottom-2">
          Copyright by 666
        </span>
      </div>
    </q-drawer>
    <q-splitter v-model="splitter" :separator-style="{ display: 'none' }">
      <template v-slot:before>
        <div
          class="q-pa-md tw-h-screen"
          style="border-right: solid 2px #f46197"
        >
          <div class="text-h4 q-mb-md text-accent">Chat History</div>
          <q-scroll-area
            style="height: 70%"
            class="tw-pr-2"
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
                v-for="history in chatHistories.messageHistory"
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
                    <span>{{ history.topic }}</span>
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
          </div>
          <span class="text-caption text-white tw-absolute tw-bottom-2">
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
            <div class="tw-w-full">
              <q-chat-message
                name="BOT"
                :avatar="botAvatar"
                bg-color="grey-1"
                text-color="white"
              >
                <div class="text-lg-body">Hi, how can i help you today ?</div>
              </q-chat-message>
              <div v-for="message in messages" :key="message.getId()">
                <q-chat-message
                  :sent="message.getStatus()"
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
                        {{ botResponse }}
                      </template>
                      <template v-else>
                        {{ message.getResponseMsg() }}
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
