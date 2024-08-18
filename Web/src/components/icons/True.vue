<template>
  <svg
    :width="size"
    :height="size"
    class="icon"
    viewBox="0 0 1280 1024"
    xmlns="http://www.w3.org/2000/svg"
  >
    <path
      d="M512 1024A512 512 0 1 1 512 0a512 512 0 0 1 0 1024z m3.008-92.992a416 416 0 1 0 0-832 416 416 0 0 0 0 832zM448 605.12L746.688 320 832 401.472 448 768 192 523.648l85.312-81.472L448 605.12z"
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
