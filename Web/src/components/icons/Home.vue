<template>
  <svg
    :width="size"
    :height="size"
    class="icon"
    viewBox="0 0 1280 1024"
    xmlns="http://www.w3.org/2000/svg"
  >
    <path
      d="M992.7 493.1L554.3 54.8c-11.6-11.6-27-16.9-42.3-16.3-15.3-0.6-30.7 4.7-42.3 16.3L31.3 493.1c-22.1 22-22.1 58.1 0 80.2 22 22.1 58.1 22.1 80.2 0L512 172.8l400.5 400.5c22.1 22.1 58.1 22.1 80.2 0 22-22.1 22-58.1 0-80.2z" fill="#515151" p-id="26443"></path><path d="M144.4 604.1V941c0 27.6 18.4 44.5 46 44.5h193V725.1c0-31.3 23.1-53.6 54.5-53.6H584c31.3 0 59.4 22.3 59.4 53.6v260.4h188.1c27.6 0 46.9-16.9 46.9-44.5V595.4l-362.6-362-371.4 370.7z"
      :fill="fill"
    />
  </svg>
</template>

<script>
export default {
  props: {
    size: {
      type: [Number, String],
      default: 64
    },
    fill: {
      type: String,
      default: "var(--el-menu-text-color)"
    }
  }
}
</script>
