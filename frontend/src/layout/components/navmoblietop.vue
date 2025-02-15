<template>
    <!-- 顶部导航栏 -->
    <div class="top-nav" v-if="isMobile">
        <div class="nav-left">
            <t-button v-if="isRootOrSecondLevelRoute" variant="text" shape="square" @click="goToUserProfile">
                <t-avatar
                    image="https://p26-passport.byteacctimg.com/img/user-avatar/424e10bd281f7bd137cd259593f07f26~140x140.awebp"
                    :hide-on-load-failed="false" />
            </t-button>
            <t-button v-else variant="text" shape="square">
                <template #icon><t-icon name="arrow-left" /></template>
            </t-button>
        </div>

        <div class="nav-center">
            <!-- <img src="/qimeng.png" alt="Logo" class="logo"> -->
        </div>
        <div class="nav-right">
            <t-button variant="text" shape="square">
                <template #icon><t-icon name="notification" /></template>
            </t-button>
        </div>
    </div>
</template>

<script setup>
import { onMounted, onUnmounted, ref, watch } from 'vue';
import { useRoute,useRouter } from 'vue-router';

// 设备类型和路由状态
const isMobile = ref(false);
const route = useRoute();
const router = useRouter();
const isRootOrSecondLevelRoute = ref(
    route.matched.length <= 2
);
// 检查屏幕宽度来更新isMobile状态
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

// 监听路由变化来更新isRootRoute状态
watch(() => route.path, () => {
    isRootOrSecondLevelRoute.value = route.matched.length <= 2;
});
const goToUserProfile = () => {

    router.push('/user');
};
</script>

<style scoped>
.top-nav {
    display: flex;
    justify-content: space-between;
    align-items: center;
    background: #fff;
    color: #fff;
    height: 60px;
    margin-bottom: 20px;
    box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
}

.nav-left {
    flex: 1;
    margin-left: 10px;
}

.nav-right {
    display: flex;
    justify-content: flex-end;
    /* 使子元素从右向左对齐 */
    align-items: center;
    /* 如果需要，使子元素在交叉轴上居中 */
    flex: 1;
    margin-right: 10px;
}

.nav-center {
    flex: 1;
    text-align: center;
}

.avatar,
.logo {
    height: 30px;
}

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
    flex-grow: 1;
}

.icon,
.icon-return {
    font-size: 24px;
    cursor: pointer;
}
</style>
