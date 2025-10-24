<template>
  <div v-if="show" class="modal-overlay" @click.self="$emit('close')">
    <div class="modal-container">
      <div class="modal-header">
        <h3>为 "{{ product.name }}" 添加视频</h3>
        <button @click="$emit('close')" class="close-btn">&times;</button>
      </div>
      <div class="modal-body">
        <textarea v-model="videoUrls" placeholder="请输入视频地址，每行一个"></textarea>
      </div>
      <div class="modal-footer">
        <button @click="$emit('close')" class="secondary-btn">取消</button>
        <button @click="save" class="primary-btn">保存</button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, watch } from 'vue';

const props = defineProps({
  show: {
    type: Boolean,
    required: true,
  },
  product: {
    type: Object,
    default: () => ({})
  }
});

const emit = defineEmits(['close', 'save']);

const videoUrls = ref('');

watch(() => props.show, (newVal) => {
  if (!newVal) {
    videoUrls.value = '';
  }
});

const save = () => {
  if (!videoUrls.value.trim()) {
    alert('请输入至少一个视频地址');
    return;
  }
  emit('save', {
    productId: props.product.ID,
    urls: videoUrls.value.split('\n').filter(url => url.trim() !== '')
  });
};
</script>

<style scoped>
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.6);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}

.modal-container {
  background-color: var(--color-background);
  border-radius: 8px;
  box-shadow: 0 5px 15px rgba(0,0,0,0.5);
  width: 90%;
  max-width: 500px;
  display: flex;
  flex-direction: column;
}

.modal-header {
  padding: 15px 20px;
  border-bottom: 1px solid var(--color-border);
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.modal-header h3 {
  margin: 0;
  color: var(--color-heading);
}

.close-btn {
  background: none;
  border: none;
  font-size: 1.5rem;
  cursor: pointer;
  color: var(--color-text);
}

.modal-body {
  padding: 20px;
}

textarea {
  width: 100%;
  min-height: 150px;
  padding: 10px;
  border: 1px solid var(--color-border);
  border-radius: 6px;
  background-color: var(--color-background-soft);
  color: var(--color-text);
  font-family: inherit;
  resize: vertical;
}

.modal-footer {
  padding: 15px 20px;
  border-top: 1px solid var(--color-border);
  display: flex;
  justify-content: flex-end;
  gap: 10px;
}

button {
  padding: 8px 16px;
  border-radius: 6px;
  cursor: pointer;
  border: 1px solid transparent;
}

.primary-btn {
  background-color: #42b983;
  color: white;
  border-color: #42b983;
}

.primary-btn:hover {
  background-color: #36a471;
}

.secondary-btn {
  background-color: var(--color-background-soft);
  color: var(--color-text);
  border-color: var(--color-border);
}

.secondary-btn:hover {
  background-color: var(--color-background-mute);
}
</style>
