<script setup>
defineProps({
  msg: {
    type: String,
    required: true,
  },
});

const login = () => {
  fetch('http://localhost:8080/login', {
    method: 'POST',
  })
    .then((res) => res.json())
    .then((data) => {
      alert(data);
    });
};

const register = () => {
  fetch('http://localhost:8080/register', {
    method: 'POST',
  })
    .then((res) => res.json())
    .then((data) => {
      alert(data);
    });
};
</script>

<template>
  <div class="greetings">
    <h1 class="red">{{ msg }}</h1>
    <div id="login-buttons">
      <button @click="login">Log In</button>
      <button @click="register">Sign Up</button>
    </div>
  </div>
</template>

<style scoped>
h1 {
  font-weight: 500;
  font-size: 2.6rem;
  position: relative;
  top: -10px;
}

#login-buttons {
  justify-content: center;
  display: flex;
  gap: 8px;
}

.greetings h1 {
  text-align: center;
}

@media (min-width: 1024px) {
  .greetings h1 {
    text-align: left;
  }
}
</style>
