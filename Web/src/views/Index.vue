<template>
  <el-container>
    <el-main>
      <el-scrollbar>
        <el-row>
          <el-col :span="12">
            <div style="display: flex; flex-direction: column; justify-items: center; align-items: center;">
              <p style="font-size: 60px; margin: 60px 0 10px 50px; color: var(--el-text-color-primary);"><b>Advanced Cybersecurity Learning Platform</b></p>
              <p style="font-size: 20px; margin: 10px 0 10px 50px; color: var(--el-text-color-primary);">MatrilxArena is an advanced cybersecurity learning platform, it uses multiple nodes to manage docker servers and provides flexible challenge, image, container management.</p>
            </div>
            <div style="margin: 20px 0 0 50px;">
              <el-button size="large"><b style="color: var(--el-text-color-primary);">Try it now</b></el-button>
              <el-button size="large"><b style="color: var(--el-text-color-primary);">Learn more about</b></el-button>
            </div>
            <div style="margin: 20px 0 0 30px;">
              <el-row>
                <el-col :span="12">
                  <el-card style="margin: 30px; border-radius: 10px; color: var(--el-text-color-primary);">
                    <p><b>Multi-node docker server management</b></p>
                  </el-card>
                </el-col>
                <el-col :span="12">
                  <el-card style="margin: 30px; border-radius: 10px; color: var(--el-text-color-primary);">
                    <p><b>Support for custom docker repositories</b></p>
                  </el-card>
                </el-col>
              </el-row>
              <el-row>
                <el-col :span="12">
                  <el-card style="margin: 30px; border-radius: 10px; color: var(--el-text-color-primary);">
                    <p><b>Flexible topic and management</b></p>
                  </el-card>
                </el-col>
                <el-col :span="12">
                  <el-card style="margin: 30px; border-radius: 10px; color: var(--el-text-color-primary);">
                    <p><b>Cross-platform software installation</b></p>
                  </el-card>
                </el-col>
              </el-row>
            </div>
          </el-col>
          <el-col :span="12">
            <div style="text-align: center;">
              <el-image src="/image/svg/matrilx.svg" style="height: 500px; margin-top: 30px"/>
              <h2 style="font-size: 50px; color: var(--el-text-color-primary); margin-top: 30px;">MatrilxArena</h2>
            </div>
          </el-col>
        </el-row>
      </el-scrollbar>
    </el-main>
  </el-container>
</template>
