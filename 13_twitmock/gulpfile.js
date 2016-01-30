var gulp = require("gulp");
var sass = require("gulp-sass");
var jade = require("gulp-jade");
var concat = require('gulp-concat');
var	autoprefixer = require('gulp-autoprefixer')

gulp.task('process-styles', function(){
	return gulp.src('src/sass/*.scss')
	.pipe(autoprefixer('last 2 versions'))
	.pipe(sass())
	.pipe(concat('style.css'))
	.pipe(gulp.dest('public/css'));
});

gulp.task('process-jade', function(){
	return gulp.src('src/jade/*.jade')
	.pipe(jade({pretty:true}))
	.pipe(gulp.dest('public/views'));
});

gulp.task('watch', function(){
	gulp.watch('src/sass/*.scss',['process-styles']);
	gulp.watch('src/jade/*.jade',['process-jade']);
});

gulp.task('default', function() {
	console.log("I have configured a gulpfile")
});