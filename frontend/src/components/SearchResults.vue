<template>
    <div class="results-window">
        <div v-for="(result, index) in results" :key="index" class="result-card">
            <div class="card-text">
                <!-- Замените 'first' и 'second' на реальные поля вашего объекта -->
                <div class="card-field">{{ result.name }}</div>
                <div class="card-field">{{ result.path }}</div>
            </div>
            <button class="upload-btn" @click="handleUpload(result)">⬇</button>
        </div>
    </div>
</template>

<script>
export default {
    name: 'SearchResults',
    props: {
        results: {
            type: Array,
            default: () => []
        }
    },
    methods: {
        handleUpload(result) {
            const downloadUrl = `http://localhost:8080/download?path=${encodeURIComponent(result.path)}`;
            window.location.href = downloadUrl;
        }
    }
}
</script>

<style scoped>
.results-window {
    width: 600px;
    height: 80vh;
    overflow-y: auto;
    border-radius: 20px;
    background-color: rgba(45, 40, 40, 0.121);
    border: 2px solid black;
    box-shadow: inset 0 0 0 2px #333;
    padding: 1rem;
}

.result-card {
    display: flex;
    justify-content: space-between;
    align-items: center;
    background-color: rgba(255, 255, 255, 0.1);
    margin-bottom: 1rem;
    padding: 0.6rem 1rem;
    border-radius: 20px;

    box-shadow: inset 0 0 0 2px #333;
}

.result-card:last-child {
    margin-bottom: 0;
}

.card-text {
    display: flex;
    flex-direction: column;
    font-family: 'Courier New', monospace;
    color: #070606;
}

.upload-btn {
    margin-left: 1rem;
    width: 40px;
    height: 40px;
    border-radius: 50%;
    background-color: rgba(8, 8, 8, 0.400);
    color: rgb(185, 55, 66);
    border: 2px solid black;
    font-family: 'Courier New', monospace;
    font-size: 18px;
    font-weight: bold;
    cursor: pointer;
    transition: background-color 0.3s ease;
}

.upload-btn:hover {
    background-color: rgba(255, 255, 255, 0.2);
}

.results-window::-webkit-scrollbar {
    width: 6px;
}

.results-window::-webkit-scrollbar-track {
    background: rgba(0, 0, 0, 0.2);
    border-radius: 10px;

    margin-top: 10px;
    margin-bottom: 10px;
}

.results-window::-webkit-scrollbar-thumb {
    background-color: rgba(185, 55, 66, 0.6);
    border-radius: 10px;
    border: 1px solid black;
}
</style>
