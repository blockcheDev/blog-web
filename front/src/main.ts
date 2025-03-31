import { createApp } from "vue";
import App from "./App.vue";
const app = createApp(App);

import "highlight.js/styles/foundation.css";

import router from "./router";
app.use(router);

import axios from "./api/axios";
app.config.globalProperties.$axios = axios;

import VMdEditor from "@kangc/v-md-editor";
import "@kangc/v-md-editor/lib/style/base-editor.css";
import githubTheme from "@kangc/v-md-editor/lib/theme/github.js";
import "@kangc/v-md-editor/lib/theme/style/github.css";
import hljs from "highlight.js";
VMdEditor.use(githubTheme, {
  Hljs: hljs,
});
app.use(VMdEditor);

app.mount("#app");
