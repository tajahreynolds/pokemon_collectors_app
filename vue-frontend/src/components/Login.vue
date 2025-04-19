<script setup>
import { useRoute, useRouter } from 'vue-router';

const router = useRouter();
const route = useRoute();

function login() {
  fetch('http://localhost:8080/login', {
    method: 'POST',
  })
    .then((res) => res.json())
    .then((data) => {
      console.log(data);
      router.push({ path: '/home' });
    });
}

function register() {
  fetch('http://localhost:8080/register', {
    method: 'POST',
  })
    .then((res) => res.json())
    .then((data) => {
      alert(data);
    });
}
</script>

<template>
  <div class="page-wrapper">
    <header>
      <div class="logo-wrapper">
        <img
          alt="Vue logo"
          class="logo"
          src="../assets/logo.svg"
          width="125"
          height="125"
        />
        <h1 class="red">Pok&#233;mon Collectors</h1>
      </div>
    </header>

    <section class="login-form">
        <h1>Let's get you signed in</h1>
        <div class="login-buttons">
          <button @click="login">Log In</button>
          <button @click="register">Sign Up</button>
        </div>
    </section>
  </div>
</template>

<style scoped>
h1 {
  font-weight: 500;
  font-size: 2.6rem;
  position: relative;
  top: -10px;
}

.login-form {
    border-radius: 16px;
    background-color: darkgrey;
    padding: 12px;
    width: fit-content;
    margin: 0px auto;
}

.login-buttons {
  justify-content: center;
  display: flex;
  gap: 8px;
}

h1 {
  text-align: center;
  white-space: nowrap;
}

@media (min-width: 1024px) {
  header {
    display: flex;
    place-items: center;
  }

  .logo {
    margin: 0 2rem 0 0;
  }

  header .logo-wrapper {
    display: flex;
    justify-content: center;
    flex-wrap: wrap;
    flex-basis: min-content;
  }

  .page-wrapper {
    justify-content: center;
    display: flex;
  }
}

header {
  line-height: 1.5;
}

.logo {
  display: block;
  margin: 0 auto 2rem;
}
</style>
