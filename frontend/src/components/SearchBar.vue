<template>
    <div class="search-bar-wrapper">
        <div class="search-container">
            <div class="search-bar">
                <input type="text" v-model="query" @keyup.enter="emitSearch" placeholder="search..." />
                <button @click="emitSearch" title="Поиск">🔍</button>
            </div>
            <div v-if="isLoading" class="progress-bar">
                <div class="progress"></div>
            </div>
        </div>

        <a href="https://github.com/FoggGhostt/Multi-functional-Search-Engine" target="_blank" rel="noopener"
            class="github-link">
            GitHub
        </a>

        <FileUploadButton @file-upload="$emit('file-upload', $event)" />
    </div>
</template>

<script>
import FileUploadButton from './FileUploadButton.vue'

export default {
    name: 'SearchComponent',
    components: { FileUploadButton },
    data() {
        return {
            query: '',
            isLoading: false
        }
    },
    methods: {
        emitSearch() {
            this.isLoading = true;
            fetch(`http://localhost:8080/api/search?query=${encodeURIComponent(this.query)}`)
                .then(res => res.json())
                .then(data => {
                    this.$emit('search-results', data)
                })
                .catch(err => console.error('Ошибка поиска:', err))
                .finally(() => {
                    this.isLoading = false;
                });
        }
    }
}
</script>

<style scoped>
.search-bar-wrapper {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
    width: 100%;
}

.search-container {
    position: relative;
}

.search-bar {
    display: flex;
    align-items: center;
    border: 2px solid black;
    box-shadow: inset 0 0 0 2px #333;
    border-radius: 999px;
    background-color: transparent;
    transition: box-shadow 0.2s ease, transform 0.2s ease;
    padding: 0.3rem 0.6rem;
}

.search-bar:hover {
    box-shadow: inset 0 0 0 2px #333, 0 0 0 2px black;
    transform: scale(1.02);
}

.search-bar input {
    flex: 1;
    width: 300px;
    height: 35px;
    border: none;
    outline: none;
    background: transparent;
    color: rgb(8, 8, 8);
    font-size: 1.1rem;
    font-family: 'Courier New', monospace;
    margin-right: 8px;
}

.search-bar input::placeholder {
    color: #0e0909;
    font-size: 14px;
}

.search-bar button {
    width: 40px;
    height: 40px;
    border-radius: 50%;
    background-color: rgba(7, 7, 7, 0.4);
    border: 2px solid black;
    color: #fff;
    cursor: pointer;
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 18px;
    transition: background-color 0.3s ease;
}

.search-bar button:hover {
    background-color: rgba(255, 255, 255, 0.2);
}

.github-link {
    position: fixed;
    bottom: 20px;
    right: 20px;
    color: rgb(185, 55, 66);
    text-decoration: none;
    font-weight: bold;
    font-family: 'Courier New', monospace;
    background-color: rgba(8, 8, 8, 0.4);
    padding: 8px 12px;
    border-radius: 999px;
    transition: background-color 0.3s ease;
    z-index: 100;
}

.github-link:hover {
    background-color: rgba(255, 255, 255, 0.2);
}

.progress-bar {
    position: absolute;
    top: 100%;
    left: 50%;
    transform: translateX(-50%);
    width: 90%; 
    margin-top: 5px;
    height: 5px;
    background: #5a4205;
    overflow: hidden;
    border-radius: 5px;
}

.progress {
    width: 100%;
    height: 100%;
    background: linear-gradient(90deg, transparent, #ddb661, transparent);
    animation: progressAnimation 2.5s infinite;
    border-radius: 5px;
}

@keyframes progressAnimation {
    0% {
        transform: translateX(-100%);
    }

    50% {
        transform: translateX(0);
    }

    100% {
        transform: translateX(100%);
    }
}
</style>