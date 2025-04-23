<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter();

function gotoHomePage() {
    router.push({ name: 'Home' });
}

const registerUsername = ref('');
const registerEmail = ref('');

function register(e) {
    e.preventDefault();
    e.stopPropagation();

    fetch('http://localhost:8080/register', {
        method: 'POST',
        body: JSON.stringify({ username: registerUsername.value, email: registerEmail.value }),
    })
        .then(gotoHomePage())
        .catch((err) => {
            console.error(err);
            alert('There was an error trying to register your account. Please try again!');
        });
}
</script>

<template>
    <div>
        <header class="flex flex-row items-center">
            <img
                alt="Vue logo"
                class="logo"
                src="../assets/logo.svg"
                width="125"
                height="125"
                @click="gotoHomePage()"
            />
            <h1 class="text-red-600">Pok&#233;mon Collectors</h1>
        </header>

        <section class="flex flex-col items-center gap-4 p-4">
            <h1 class="text-center min-[551px]:max-[976px]:max-w-[532px]">
                Sign up for a free account to gain access to all features!
            </h1>
            <form
                class="w-3/4 min-w-[277px] rounded-lg bg-white p-4 min-[610px]:max-[786px]:min-w-[523px]"
                @submit="register"
            >
                <div class="pb-4">
                    <label
                        for="username"
                        class="text-dark block max-[786px]:min-[610px]:whitespace-nowrap"
                        >Username (this is how you will be referred to across the site):</label
                    >
                    <input
                        type="text"
                        id="username"
                        name="username"
                        class="text-dark w-full rounded border px-3 py-2"
                        placeholder="Choose your username"
                        v-model="registerUsername"
                        required
                    />
                </div>
                <div class="pb-4">
                    <label
                        for="email"
                        class="text-dark block"
                        >Email Address (we will send you a link to complete your
                        registration):</label
                    >
                    <input
                        type="email"
                        id="email"
                        name="email"
                        class="w-full rounded border px-3 py-2 text-gray-700"
                        placeholder="Enter your email"
                        v-model="registerEmail"
                        required
                    />
                </div>
                <div class="text-dark flex items-center justify-between">
                    <span>Already have an account? <RouterLink to="Login">Login</RouterLink></span>
                    <button
                        class="btn-red mr-4 text-nowrap"
                        type="submit"
                    >
                        Register Now
                    </button>
                </div>
            </form>
        </section>
    </div>
</template>

<style scoped></style>
