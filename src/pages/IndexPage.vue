<script setup lang="ts">
import { Message } from "src/constants/message";
import { MessageInterface, History } from "src/constants";
import { ref, Ref, watch, reactive, onMounted, computed } from "vue";
import { QScrollArea, useQuasar } from "quasar";
import { dummyResponse, allHistory } from "src/constants/history";
import { useUtility } from "src/composables/useUtility";
import { useMessages } from "src/composables/useMessages";
import { botAvatar, userAvatar } from "src/constants/avatar";

const $q = useQuasar();
const BREAKPOINT = 800;
const isSmallScreen = computed(() => $q.screen.width < BREAKPOINT);

const splitter = computed(() => (isSmallScreen.value ? 0 : 20));
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
// TODO: chat history
// This must be updated to the newest one (history length + 1)
const currentConversationID = ref(0);

// Perform fetching data to fill this array
const messages: MessageInterface[] = reactive([]);
const userInput: Ref<string> = ref("");
const botResponse: Ref<string> = ref("");
const botFullResponse: Ref<string> = ref("");
const isResponding: Ref<boolean> = ref(false);

const { animateMessage, random } = useUtility({
  startNum: 1,
  endNum: 6,
  botMessage: botResponse,
  duration: 20,
});

const method = ref("");

const newChat = () => {
  currentConversationID.value = allHistory.messageHistory.length + 1;
  messages.splice(0, messages.length);
};

const sendMessage = () => {
  if (!isResponding.value) {
    const { generateMessageId } = useMessages({ allHistory });
    isResponding.value = true;
    scrollToBottom();
    let userMessage: Message;
    if (messages.length === 0) {
      const availId = generateMessageId();
      userMessage = new Message(
        messages.length + 1,
        true,
        userInput.value,
        new Date().toLocaleTimeString(),
        availId
      );
    } else {
      userMessage = new Message(
        messages.length + 1,
        true,
        userInput.value,
        new Date().toLocaleTimeString(),
        currentConversationID.value
      );
    }

    // Send request to backend
    messages.push(userMessage);
    const currentTopic = userInput.value;
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
      allHistory.messageHistory.push(newHistory);
    }
  }
};

watch(messages, () => {
  scrollToBottom();
});

onMounted(() => scrollToBottom());

const switchConversation = () => {
  // replace the current array to chosen history conversation
  const chosenHistory =
    allHistory.messageHistory[currentConversationID.value - 1];
  messages.splice(0, messages.length, ...chosenHistory.conversation);
};
</script>
<template>
  <q-page class="tw-h-screen">
    <q-splitter v-model="splitter" :separator-style="{ display: 'none' }">
      <template v-slot:before>
        <div
          class="q-pa-md tw-h-screen"
          style="border-right: solid 2px #000000"
        >
          <div class="text-h4 q-mb-md">Chat History</div>
          <q-scroll-area style="height: 70%" class="tw-pr-2">
            <q-btn
              label="New Chat"
              icon="add_circle"
              class="tw-w-full tw-h-12"
              @click="newChat"
              :disable="isResponding"
            />

            <q-tabs
              v-model="currentConversationID"
              vertical
              class="tw-py-4 text-body"
              active-bg-color="primary"
              active-color="white"
            >
              <template
                v-for="history in allHistory.messageHistory"
                :key="history"
              >
                <q-tab
                  :name="history.historyId"
                  @click="switchConversation"
                  style="border-bottom: 2px solid #000000"
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
          <div class="tw-mt-10">
            <h3>Algorithm:</h3>
            <q-radio
              v-model="method"
              checked-icon="task_alt"
              unchecked-icon="panorama_fish_eye"
              val="KMP"
              label="KMP"
            />
            <q-radio
              v-model="method"
              checked-icon="task_alt"
              unchecked-icon="panorama_fish_eye"
              val="BoyerMoore"
              label="Boyer Moore"
            />
          </div>
        </div>
      </template>

      <template v-slot:after>
        <div
          style="width: 100%"
          class="tw-h-screen tw-w-flex tw-flex-col tw-justify-end"
        >
          <q-scroll-area
            style="height: 90%"
            :delay="1000"
            class="tw-px-5 chat-area"
            ref="scrollArea"
          >
            <div class="tw-w-full">
              <q-chat-message name="BOT" :avatar="botAvatar" bg-color="amber">
                <div class="text-lg-body">Hi, how can i help you today ?</div>
              </q-chat-message>
              <div v-for="message in messages" :key="message.getId()">
                <q-chat-message
                  :sent="message.getStatus()"
                  text-color="white"
                  bg-color="primary"
                  :avatar="userAvatar"
                  :stamp="message.getSentTime()"
                >
                  <div class="text-lg-body">{{ message.getText() }}</div>
                </q-chat-message>
                <q-chat-message name="BOT" :avatar="botAvatar" bg-color="amber">
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
          <div class="row items-center tw-gap-x-2 tw-px-2 tw-mt-2">
            <q-input
              v-model="userInput"
              label="Type your message here"
              class="tw-grow"
              @keyup.enter="sendMessage"
              autogrow
              outlined
              rounded
            />
            <div>
              <q-btn
                icon="send"
                @click="sendMessage"
                size="md"
                round
                color="primary"
                :disable="
                  isResponding ||
                  userInput.length === 0 ||
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

:deep(.q-btn__content) {
  font-size: 16px;
  text-transform: none !important;
}

:deep(.q-field__control-container) {
  padding-inline: 10px;
}
:deep(.q-field__label) {
  margin-left: 10px !important;
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
