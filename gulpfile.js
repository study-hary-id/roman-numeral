const gulp = require("gulp");
const shell = require("gulp-shell");

const files = ["*.go", "*.mod"];

// This compiles new binary with source change.
gulp.task("install-binary", shell.task([
    "go install ."
]));

// Second argument tells install-binary is a dependency for this task.
gulp.task(
    "restart-supervisor",
    shell.task(["supervisorctl restart roman-numeral-api"])
);

gulp.task("watch", function () {
    // Watch the source code for all changes.
    gulp.watch(files, gulp.series("install-binary", "restart-supervisor"));
});

gulp.task("default", gulp.series("watch"));
