<template>
  <div id="app">
    <div class="row">
      <div class="left">
        <textarea v-model="input"></textarea>
      </div>
      <div class="right">
        <div class="content">{{ result }}</div>
      </div>
    </div>
    <div class="actions">
      <button @click="hash">Sha256 Hash</button>
      <button @click="encode">Base64 Encode</button>
    </div>
  </div>
</template >

<script>
import axios from "axios";
const API_BASE_URL = "http://localhost/api";

export default {
  data() {
    return {
      input: "",
      result: "",
    };
  },

  methods: {
    async hash() {
      const api = `${API_BASE_URL}/hash/sha256`;
      const body = { input: this.input };
      const response = await axios.post(api, body);
      if (response.status === 200) {
        this.result = response.data.result;
      }
    },

    async encode() {
      const api = `${API_BASE_URL}/encode/base64`;
      const body = { input: this.input };
      const response = await axios.post(api, body);
      if (response.status === 200) {
        this.result = response.data.result;
      }
    },
  },
};
</script>

<style scoped>
#app {
  font-family: "IBM Plex Mono";
}

.row {
  display: flex;
  flex-direction: row;
  align-items: stretch;
  justify-content: space-between;
}

.row > div {
  flex: 0 0 50%;
  max-width: 50%;
}

.content {
  height: 300px;
  text-align: left;
  word-wrap: break-word;
  background-color: #a4f2d8;
}

.left {
  display: flex;
  flex-direction: column;
}

.left > textarea {
  flex: 1;
}

div.actions {
  margin-top: 10px;
  display: flex;
  justify-content: center;
}
</style>
