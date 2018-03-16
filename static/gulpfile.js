var gulp = require("gulp");
var source = require("vinyl-source-stream");
var babelify = require("babelify");
var browserify = require("browserify");
var template = require("gulp-template");

// function time() {
//     return "[" + new Date().toLocaleTimeString() + "] ";
// }
//
// var bundlers = [
//     "./js/index.jsx"
// ].map(function(entry) {
//     return browserify({
//         entries: [entry],
//         transform: [babelify,{presets: ["es2015", "react"]}],
//         debug: true
//     });
// });
//
// function build(bundler) {
//     (bundler ? [bundler] : bundlers).map(function(bundler) {
//         var entry = bundler._options.entries[0];
//         var target = entry.replace(/^.*\/([^\/]+js)x?$/, "$1");
//         var start = Date.now();
//         console.log(time() + "Bundling " + entry + " ...");
//         bundler.bundle()
//             .on("error", console.error)
//             .on("end", function() {
//                 console.log(time() + "Finished " + target + " after " + (Date.now() - start) + " ms")
//             })
//             .pipe(source(target))
//             .pipe(gulp.dest("dist"));
//     });
//
//     // gulp.src(cssfiles)
//     //     .pipe(concat('app.min.css'))
//     //     .pipe(gulp.dest("dist"));
//
//     gulp.src("html/index.html")
//         .pipe(template({
//             ts: Date.now()
//         }))
//         .pipe(gulp.dest("../views"));
// }
//
// gulp.task("watch", function() {
//     bundlers = bundlers.map(function(bundler) {
//         return watchify(bundler);
//     });
//
//     bundlers.forEach(function(bundler) {
//         bundler.on("update", function() {
//             build(bundler)
//         });
//     });
//
//     build();
// });
//
// gulp.task("build", function() {
//     build()
// });
// gulp.task("default", ["build"]);

gulp.task('default', function(){
  browserify("js/index.jsx")
  .transform("babelify",{presets: ["es2015", "react"]})
  .bundle()
  .pipe(source("index.js"))
  .pipe(gulp.dest("dist"));

  gulp.src("html/index.html")
      .pipe(gulp.dest("../views"));
})
