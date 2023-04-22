<script setup lang="ts">
import { Message } from "src/constants/message";
import { ref, Ref, watch, reactive, onMounted } from "vue";
import { QScrollArea } from "quasar";
import { dummyResponse } from "src/constants/history";
import { useUtility } from "src/composables/useUtility";
import { botAvatar, userAvatar } from "src/constants/avatar";

const splitter = ref(20);
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

// Perform fetching data to fill this array
const messages: Message[] = reactive([]);
const userInput: Ref<string> = ref("");
const botResponse: Ref<string> = ref("");
const botFullResponse: Ref<string> = ref("");

const { animateMessage, random } = useUtility({
  startNum: 1,
  endNum: 6,
  botMessage: botResponse,
  duration: 20,
});

const method = ref("");

const sendMessage = () => {
  scrollToBottom();
  const userMessage = new Message(
    messages.length + 1,
    true,
    userInput.value,
    new Date().toLocaleTimeString()
  );

  // Send request to backend
  messages.push(userMessage);
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
  }, 1500);
};

watch(messages, () => {
  scrollToBottom();
});

onMounted(() => scrollToBottom());
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
            <div
              v-for="n in 20"
              :key="n"
              class="q-my-md tw-py-4 text-body"
              style="border-bottom: 2px solid #000000"
            >
              Chat topic - {{ n }}
            </div>
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
              <transition name="typing">
                <q-chat-message name="BOT" :avatar="botAvatar" bg-color="amber">
                  <div class="text-lg-body">Hi, how can i help you today ?</div>
                </q-chat-message>
              </transition>
              <div v-for="message in messages" :key="message.getId()">
                <q-chat-message
                  :sent="message.getStatus()"
                  :label="message.getSentTime()"
                  text-color="white"
                  bg-color="primary"
                  :avatar="userAvatar"
                >
                  <div class="text-lg-body">{{ message.getText() }}</div>
                </q-chat-message>

                <q-chat-message name="BOT" :avatar="botAvatar" bg-color="amber">
                  <template v-if="message.getResponseCode() === 0">
                    <q-spinner-dots size="2rem" />
                  </template>
                  <template v-else>
                    <transition name="typing">
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
                    </transition>
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
    align-items: end !important;
    position: relative;
  }
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
.typing-enter-active {
  animation: typing 2s;
}
@keyframes typing {
  from {
    width: 0;
  }
  to {
    width: 100%;
  }
}
</style>
