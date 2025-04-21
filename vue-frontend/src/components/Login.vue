<script setup>
import { useRouter } from 'vue-router';

const router = useRouter();

function login() {
    fetch('http://localhost:8080/login', {
        method: 'POST',
    })
        .then((res) => res.json())
        .then((data) => {
            console.log(data);
            router.push({ name: 'Home' });
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
                <h1 class="text-red-600">Pok&#233;mon Collectors</h1>
            </div>
        </header>

        <section class="login-form bg-gray-200">
            <h1>Let's get you signed in</h1>
            <div class="flex flex-col justify-center gap-y-2">
                <button
                    class="btn-red"
                    @click="login"
                >
                    Log In
                </button>
                <button
                    class="btn-red"
                    @click="register"
                >
                    Sign Up
                </button>
                <button class="btn-white">Sign In With Google</button>
            </div>
        </section>
    </div>
</template>

<style scoped>
header {
    line-height: 1.5;
}

h1 {
    font-weight: 500;
    font-size: 2.6rem;
    position: relative;
    top: -10px;
    text-align: center;
    white-space: nowrap;
}

.logo {
    display: block;
    margin: 0 auto 2rem;
}

.login-form {
    color: var(--color-background);
    border-radius: 16px;
    padding: 12px;
    width: fit-content;
    margin: 0px auto;
}

@media (min-width: 1024px) {
    header {
        display: flex;
        place-items: center;
    }

    .logo {
        margin: 0 2rem 0 0;
    }

    .logo-wrapper {
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
</style>
