<template>
    <div class="upload-wrapper">
        <input type="file" ref="fileInput" @change="handleFileChange" multiple hidden />
        <button :class="['upload-button', uploadStatus]" @click="triggerFileInput" title="Загрузить файл">
            ⬆
        </button>
    </div>
</template>

<script>
import axios from 'axios';

export default {
    data() {
        return {
            uploadStatus: null
        }
    },
    methods: {
        triggerFileInput() {
            this.$refs.fileInput.click();
        },
        handleFileChange(event) {
            const files = event.target.files;
            if (files.length > 0) {
                const formData = new FormData();
                for (let i = 0; i < files.length; i++) {
                    formData.append('files', files[i]);
                }
                axios.post('http://localhost:8080/upload', formData, {
                    headers: { 'Content-Type': 'multipart/form-data' }
                })
                    .then(response => {
                        console.log('Файлы успешно загружены', response.data);
                        this.uploadStatus = 'success';
                        this.$refs.fileInput.value = null;
                        setTimeout(() => this.uploadStatus = null, 1000);
                    })
                    .catch(error => {
                        console.error('Ошибка загрузки файлов', error);
                        this.uploadStatus = 'error';
                        this.$refs.fileInput.value = null;
                        setTimeout(() => this.uploadStatus = null, 1000);
                    });
            }
        }
    }
}
</script>

<style scoped>
.upload-wrapper {
    margin-left: 12px;
}

.upload-button {
    width: 40px;
    height: 40px;
    border-radius: 50%;
    background-color: rgba(7, 7, 7, 0.4);
    color: rgb(170, 55, 55);
    border: 2px solid black;
    font-family: 'Courier New', monospace;
    font-size: 18px;
    font-weight: bold;
    cursor: pointer;
    transition: background-color 0.3s ease;
}

.upload-button:hover {
    background-color: rgba(255, 255, 255, 0.2);
}

.upload-button.success {
    background-color: green !important;
}

.upload-button.error {
    background-color: red !important;
}
</style>