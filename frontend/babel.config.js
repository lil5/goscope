module.exports = {
  presets: ["@vue/cli-plugin-babel/preset"],
  plugins: [
    [
      "prismjs",
      {
        languages: [
          "javascript",
          "css",
          "markup",
          "go",
          "typescript",
          "java",
          "groovy",
          "sql",
          "shell-session",
          "shell",
          "json"
        ],
        plugins: [],
        theme: "none",
        css: false
      }
    ]
  ]
};
