<script setup lang="ts">
import { Message } from 'src/constants/message';
import { ref, Ref, watch, reactive, onMounted } from 'vue';
import { QScrollArea } from 'quasar';

const splitter = ref(20);
const scrollArea = ref<QScrollArea | null>(null);
const scrollToBottom = () => {
  const scrollTarget = scrollArea.value?.getScrollTarget();
  const scrollPosition = scrollArea.value?.getScrollPosition();
  if (scrollTarget && scrollPosition) {
    scrollArea.value?.setScrollPosition(
      'vertical',
      scrollTarget.scrollHeight,
      1000
    );
  }
};

const messages: Message[] = reactive([]);
const userInput: Ref<string> = ref('');
const sendMessage = () => {
  scrollToBottom();
  const userMessage = new Message(
    messages.length + 1,
    true,
    userInput.value,
    new Date().toLocaleTimeString()
  );
  messages.push(userMessage);
  userInput.value = '';
  setTimeout(() => {
    messages[messages.length - 1].setResponse('Reponse from Bot', 200);
  }, 1500);
};

watch(messages, () => {
  scrollToBottom();
});

onMounted(() => scrollToBottom());
</script>
<template>
  <q-page class="row items-center justify-evenly">
    <q-splitter v-model="splitter">
      <template v-slot:before>
        <div class="q-pa-md">
          <div class="text-h4 q-mb-md">Before</div>
          <div v-for="n in 5" :key="n" class="q-my-md">
            {{ n }}. Lorem ipsum dolor sit, amet consectetur adipisicing elit.
            Quis praesentium cumque magnam odio iure quidem, quod illum numquam
            possimus obcaecati commodi minima assumenda consectetur culpa fuga
            nulla ullam. In, libero.
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
            class="tw-px-5"
            ref="scrollArea"
          >
            <div class="tw-w-full">
              <q-chat-message
                name="BOT"
                avatar="https://cdn.quasar.dev/img/avatar5.jpg"
                bg-color="amber"
              >
                <div>Hi, how can i help you today ?</div>
              </q-chat-message>
              <div v-for="message in messages" :key="message.getId()">
                <q-chat-message
                  :sent="message.getStatus()"
                  :label="message.getSentTime()"
                  text-color="white"
                  bg-color="primary"
                  avatar="https://cdn.quasar.dev/img/avatar.png"
                >
                  <div>{{ message.getText() }}</div>
                </q-chat-message>

                <q-chat-message
                  name="BOT"
                  avatar="https://cdn.quasar.dev/img/avatar5.jpg"
                  bg-color="amber"
                >
                  <template v-if="message.getResponseCode() === 0">
                    <q-spinner-dots size="2rem" />
                  </template>
                  <template v-else>
                    <div>
                      {{ message.getResponseMsg() }}
                    </div>
                  </template>
                </q-chat-message>
              </div>
            </div>
          </q-scroll-area>
          <div class="row">
            <q-input
              v-model="userInput"
              label="Type your message here"
              class="tw-grow"
              autogrow
              outlined
              rounded
            />
            <q-btn icon="send" @click="sendMessage" size="md" flat unelevated />
          </div>
        </div>
      </template>
    </q-splitter>
  </q-page>
</template>

<style lang="scss" scoped>
:deep(.q-scrollarea__content) {
  display: flex;
  width: 100% !important;
  align-items: end !important;
  position: relative;
}
</style>
