import { getView } from "../lib/lib.js";
const host = location.host;

const baseUrl = "http://" + host;

export async function get(path) {
  const url = baseUrl + path;
  const token = localStorage.getItem("auth-token") || "none";

  try {
    const response = await fetch(url, {
      method: "GET",
      headers: {
        Accept: "application/json",
        "auth-token": JSON.stringify({ token }),
      },
    });
    const data = await response.json();

    if (response.ok) {
      return data;
    } else dispatchEvent(new CustomEvent("errorOccur", { detail: data }));
  } catch (error) {
    console.log(error);
  }
}

export async function post(path, data) {
  const url = baseUrl + path;
  const token = localStorage.getItem("auth-token") || "none";

  try {
    const response = await fetch(url, {
      method: "POST",
      body: JSON.stringify(data),
      headers: {
        "auth-token": JSON.stringify({ token }),
      },
    });
    const dataObject = await response.json();

    if (response.ok) {
      return dataObject;
    } else dispatchEvent(new CustomEvent("errorOccur", { detail: dataObject }));
  } catch (error) {
    console.log(error);
  }
}

export async function del(path) {
  const url = baseUrl + path;
  const token = localStorage.getItem("auth-token") || "none";

  return fetch(url, {
    method: "DELETE",
    headers: {
      Accept: "application/json",
      "auth-token": JSON.stringify({ token }),
    },
  })
    .then((response) => {
      return response.json();
    })
    .then((data) => data);
}

export async function logout() {
  const path = "/auth/signout";
  return del(path);
}

export async function getPosts(query) {
  const path = `/posts/get?${query}`;
  return await get(path);
}

export async function getPostComments(queries) {
  const path = `/post/comments?${queries}`;
  return await get(path).catch((error) => error);
}

export async function postComment(postId, data) {
  const path = `/post/comment/create?postid=${postId}`;
  return post(path, data).catch((error) => error);
}

export async function createPost(data) {
  const path = "/post/create";
  return await post(path, data).catch((error) => error);
}

export async function checkSession() {
  const token = localStorage.getItem("auth-token");
  return await post("/verifsess", { token });
}

export function setWSConnection() {
  const token = localStorage.getItem("auth-token");
  if (window.WebSocket) {
    const socket = new WebSocket(
      `ws://${host}/ws?token=${JSON.stringify({ token })}}`
    );
    window.app.wsConnection = socket;
  }
}

export async function getChatByUser(username) {
  const path = `/chat?username=${username}`;
  return await get(path);
}

export async function getUsersStatus() {
  const path = "/users-status";
  return await get(path);
}

export const EventType = {
  WS_JOIN_EVENT: "join-event",
  WS_DISCONNECT_EVENT: "disconnect-event",
  WS_MESSAGE_EVENT: "msg-event",
  WS_READ_EVENT: "read-event",
  WS_SEND_EVENT: "send-event",
  WS_REACT_EVENT: "react-event",
  WS_NEW_POST_EVENT: "new-post-event",
  WS_NEW_USER_EVENT: "new-user-event",
  WS_TYPING_EVENT: "typing-event",
};

export async function getMessages(queries) {
  const path = `/messages?${queries}`;
  return await get(path);
}

export async function postReact(postId, react) {
  const path = `/post/react?postid=${postId}&react=${react}`;
  return get(path);
}

export function handleWSEvent(wsEvent) {
  const event = JSON.parse(wsEvent);
  switch (event["Type"]) {
    case EventType.WS_NEW_USER_EVENT:
      dispatchEvent(new CustomEvent("newUser", { detail: event["Data"] }));
      break;
    case EventType.WS_NEW_POST_EVENT:
      dispatchEvent(new CustomEvent("newPost", { detail: event["Data"] }));
      break;
    case EventType.WS_JOIN_EVENT:
      setStatusOnline(event.From);
      break;
    case EventType.WS_DISCONNECT_EVENT:
      setStatusOffline(event.From);
      break;
    case EventType.WS_MESSAGE_EVENT:
      dispatchEvent(new CustomEvent("newMessage", { detail: event }));
      if (app.user.username != event.From) {
        const statusItmView = getView(`user-status${event.From}`);
        statusItmView.user.unread_count += 1;
        dispatchEvent(
          new CustomEvent("unreadMsg", { detail: statusItmView.user })
        );
      }
      break;
    case EventType.WS_TYPING_EVENT:
      dispatchEvent(new CustomEvent("typing", { detail: event }));
      break;
    default:
      break;
  }
}

function setStatusOnline(username) {
  const userList = getView("userList");

  if (!userList.refreshing) {
    // try{
    const dot = getView(`${username}status-dot`).element;
    const text = getView(`${username}-status-text`).element;
    const statusItmView = getView(`user-status${username}`);
    statusItmView.user.status = "online";
    text.innerText = "online";
    dot.style.backgroundColor = "green";
    if (getView(`chat${username}`)) {
      const dot = getView(`chat-${username}-status-dot`).element;
      dot.style.backgroundColor = "green";
    }
    // }catch{}
  }
}

function setStatusOffline(username) {
  const userList = getView("userList");

  if (!userList.refreshing) {
    // try {
    const dot = getView(`${username}status-dot`).element;
    const text = getView(`${username}-status-text`).element;
    const statusItmView = getView(`user-status${username}`);
    statusItmView.user.status = "offline";
    text.innerText = "offline";
    dot.style.backgroundColor = "gray";
    if (getView(`chat${username}`)) {
      const dot = getView(`chat-${username}-status-dot`).element;
      dot.style.backgroundColor = "gray";
    }
    // } catch { }
  }
}
