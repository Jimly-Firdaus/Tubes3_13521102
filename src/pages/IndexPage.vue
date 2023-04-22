<template>
  <q-page class="row items-center justify-evenly">
    <q-splitter v-model="splitter">
      <template v-slot:before>
        <div class="q-pa-md">
          <div class="text-h4 q-mb-md">Before</div>
          <div v-for="n in 20" :key="n" class="q-my-md">
            {{ n }}. Lorem ipsum dolor sit, amet consectetur adipisicing elit.
            Quis praesentium cumque magnam odio iure quidem, quod illum numquam
            possimus obcaecati commodi minima assumenda consectetur culpa fuga
            nulla ullam. In, libero.
          </div>
        </div>
      </template>

      <template v-slot:after>
        <div style="width: 100%" class="bg-red tw-h-screen">
          <q-chat-message
              v-if="isFetchingReponse"
              name="BOT"
              avatar="https://cdn.quasar.dev/img/avatar5.jpg"
              bg-color="amber"
            >
              <div>
                Hi, how can i help you today ?
              </div>
            </q-chat-message>
          <div v-for="message in messages" :key="message.getId()">
            <q-chat-message
              :sent="message.getStatus()"
              label="John Doe"
              stamp="4 minutes ago"
              text-color="white"
              bg-color="primary"
            >
              <q-avatar>
                <img src="https://cdn.quasar.dev/img/avatar.png" />
              </q-avatar>
              <div>{{ message.getText() }}</div>
            </q-chat-message>

            <q-chat-message
              v-if="isFetchingReponse"
              name="BOT"
              avatar="https://cdn.quasar.dev/img/avatar5.jpg"
              bg-color="amber"
            >
              <q-spinner-dots size="2rem" />
            </q-chat-message>
          </div>

          <q-input v-model="userInput" label="Type your message here" />
          <q-btn label="Send" @click="sendMessage" />
        </div>
      </template>
    </q-splitter>
  </q-page>
</template>

<script setup lang="ts">
import { Message } from 'src/constants/message';
import { ref, Ref, watch } from 'vue';

const splitter = ref(20);

const messages: Ref<Message[]> = ref([]);
const userInput: Ref<string> = ref('');
const isFetchingReponse: Ref<boolean> = ref(true);
const sendMessage = () => {
  const userMessage = new Message(
    messages.value.length + 1,
    true,
    userInput.value
  );
  setTimeout(() => {
    isFetchingReponse.value = false;
  })
  messages.value.push(userMessage);
  userInput.value = '';
};
watch(userInput, () => {
  isFetchingReponse.value = true;
})

</script>
