<template>
    <div class="footer-container">
        <div class="footer-content">
            <a href="https://beian.miit.gov.cn/" target="_blank" rel="noopener noreferrer">
                粤ICP备2023024319号-1
            </a>
        </div>
    </div>
</template>

<style scoped>
.footer-container {
    width: 100%;
    margin-top: 2rem;
    padding: 1rem 0;
    text-align: center;
}

.footer-content {
    display: inline-block;
    padding: 0.5rem 1rem;
    border-radius: 4px;
    background-color: #eeeeee;
}

.footer-content a {
    color: #666;
    font-size: 0.875rem;
    text-decoration: none;
    transition: color 0.2s;
}

.footer-content a:hover {
    color: #1890ff;
}

@media (prefers-color-scheme: dark) {
    .footer-content {
        background-color: rgba(60, 60, 60, 0.8);
    }

    .footer-content a {
        color: #ccc;
    }
}
</style>
