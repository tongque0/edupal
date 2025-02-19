<template>
    <div class="layout-container">
        <header class="header-fixed">
            <div class="nav-container">
                <NavPC />
                <Navmobile />
            </div>
        </header>

        <main class="main-content">
            <div class="content-wrapper">
                <slot></slot>
            </div>
        </main>
    </div>
</template>

<script setup>
import NavPC from '@/layout/components/navpc.vue';
import Navmobile from '@/layout/components/navmoblie.vue';
</script>

<style scoped>
.layout-container {
    display: flex;
    flex-direction: column;
    min-height: 100vh;
    width: 100%;
    overflow-x: hidden; /* 防止横向滚动 */
}

.header-fixed {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    z-index: 100;
    background: #fff;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.08);
}

.nav-container {
    max-width: 1200px;
    margin: 0 auto;
    width: 100%;
}

.main-content {
    margin-top: 64px;
    flex: 1;
    width: 100%;
    display: flex;
    justify-content: center;
}

.content-wrapper {
    width: 100%;
    max-width: 1200px;
    margin: 0 auto;
    padding: 20px 16px;
    box-sizing: border-box; /* 确保padding不会增加总宽度 */
}

/* 移动端适配 */
@media screen and (max-width: 768px) {
    .main-content {
        margin-top: 56px;
    }

    .content-wrapper {
        padding: 12px;
        max-width: 100%; /* 移动端下不限制最大宽度 */
    }

    .nav-container {
        padding: 0 12px;
        max-width: 100%;
    }
}
</style>
