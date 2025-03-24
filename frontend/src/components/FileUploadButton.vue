<template>
    <div class="upload-wrapper">
        <input type="file" ref="fileInput" @change="handleFileChange" hidden />
        <button class="upload-button" @click="triggerFileInput" title="Загрузить файл">
            ⬆
        </button>
    </div>
</template>

<script>
export default {
    methods: {
        triggerFileInput() {
            this.$refs.fileInput.click();
        },
        handleFileChange(event) {
            const file = event.target.files[0];
            if (file) {
                this.$emit('file-upload', file);
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
    background-color: rgba(7, 7, 7, 0.400);
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
</style>