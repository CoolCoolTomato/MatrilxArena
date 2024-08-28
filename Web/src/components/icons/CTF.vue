<template>
  <svg
    :width="size"
    :height="size"
    class="icon"
    viewBox="0 0 1280 1024"
    xmlns="http://www.w3.org/2000/svg"
  >
    <path
      d="M153.6 554.688V1024H64V0h358.656l183.488 85.312H960V640H601.344L426.432 554.688H153.6z"
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
