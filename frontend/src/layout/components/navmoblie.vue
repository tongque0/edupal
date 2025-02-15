<template>
    <div v-if="isMobile" class="nav-bar">
        <div class="nav-item" v-for="item in list" :key="item.value">
            <t-button variant="text" shape="square" @click="changeHandler(item.value)">
                <template #icon><t-icon :name="item.icon" /></template>
            </t-button>
        </div>
    </div>
</template>

<script setup>
import { onMounted, onUnmounted, ref } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter();
const list = ref([
    { value: 'home', icon: 'home', ariaLabel: '首页' },
    { value: 'user', icon: 'app', ariaLabel: '软件' },
    { value: 'user', icon: 'chat', ariaLabel: '聊天' },
    { value: 'user', icon: 'user', ariaLabel: '我的' },
]);
const isMobile = ref(false);
const changeHandler = (active) => {
    console.log('change', active);
    // 根据active值进行路由跳转
    if (active === 'home') {
        router.push('/');
    } else if (active === 'user') {
        router.push('/user');
    }
};
// 定义函数来检查屏幕宽度
const checkScreenWidth = () => {
    isMobile.value = window.innerWidth < 768;
};

onMounted(() => {
    checkScreenWidth();
    window.addEventListener('resize', checkScreenWidth);
});

onUnmounted(() => {
    window.removeEventListener('resize', checkScreenWidth);
});
</script>

<style scoped>
.nav-bar {
    display: flex;
    justify-content: space-around;
    background: #fff;
    padding: 10px 0;
    position: fixed;
    bottom: 0;
    left: 0;
    right: 0;
    height: 35px;
    box-shadow: 0 -2px 5px rgba(0, 0, 0, 0.1);
}

.nav-item {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    /* 确保沿主轴居中 */
    flex-grow: 1;
}

.icon {
    font-size: 30px;
}
</style>
