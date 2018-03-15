var gulp = require("gulp");
var source = require("vinyl-source-stream");
var babelify = require("babelify");
var browserify = require("browserify");

gulp.task('default', function(){
  browserify("js/index.jsx")
  .transform("babelify",{presets: ["es2015", "react"]})
  .bundle()
  .pipe(source("index.js"))
  .pipe(gulp.dest("dist"));

  gulp.src("html/index.html")
      .pipe(gulp.dest("../views"));
})
