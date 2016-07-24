var gulp = require('gulp');
var plumber = require('gulp-plumber');
var browserSync = require('browser-sync');
var inject = require('gulp-inject');
var reload = browserSync.reload;


gulp.task('default', ['print']);


gulp.task('build', function(){

  //Copy App File
  gulp.src(['src/**/', '!src/js/*.js', '!src/css/*.css', '!src/*.html'])
      .pipe(gulp.dest('./build/'));

  //Copy Minifier File To Src Floder
  gulp.src(['bower_components/angular/angular.min.js',
            'bower_components/jquery/dist/jquery.min.js',
            'bower_components/semantic/dist/semantic.min.js',
            'bower_components/angular-ui-router/release/angular-ui-router.min.js',
            'bower_components/angular-websocket/dist/angular-websocket.min.js'])
      .pipe(gulp.dest('./build/js/'));

  gulp.src(['bower_components/semantic/dist/semantic.min.css'])
      .pipe(gulp.dest('./build/css/'));

  var fileOrder = [
    './build/js/angular.min.js',
    './build/js/angular-ui-router.min.js',
    './build/js/angular-websocket.min.js',
    './build/**/*.js',
    './build/**/*.css'
  ];

  //Inject Script
  gulp.src('src/*.html')
      .pipe(gulp.dest('./build/'))
      .pipe(inject(gulp.src(fileOrder, {read: false}), {relative: true}))
      .pipe(gulp.dest('build'));

});


gulp.task('build:dev', function(){
  //Copy Minifier File To Src Floder
  gulp.src(['bower_components/angular/angular.js',
            'bower_components/jquery/dist/jquery.js',
            'bower_components/semantic/dist/semantic.js',
            'bower_components/angular-ui-router/release/angular-ui-router.js',
            'bower_components/angular-websocket/dist/angular-websocket.js'])
      .pipe(gulp.dest('./src/js/'));

  gulp.src(['bower_components/semantic/dist/semantic.css'])
      .pipe(gulp.dest('./src/css/'));

  var fileOrder = [
    './src/js/angular.js',
    './src/js/angular-ui-router.js',
    './src/js/angular-websocket.js',
    './src/**/*.js',
    './src/**/*.css'
  ];

  gulp.src('src/*.html')
      .pipe(inject(gulp.src(fileOrder, {read: false}), {relative: true}))
      .pipe(gulp.dest('src'));
});


gulp.task('serve', ['build'], function () {
    browserSync({
        notify: false,
        port: 9000,
        server: {
            baseDir: ['build']
        }
    });

    gulp.watch([
        'build/**/*.html',
        'build/**/*.js',
        'build/**/*.css'
    ], ['build']).on('change', reload);

});

gulp.task('serve:dev', ['build:dev'], function () {
    browserSync({
        notify: false,
        port: 9000,
        server: {
            baseDir: ['src']
        }
    });

    gulp.watch([
        'src/**/*.html',
        'src/**/*.js',
        'src/**/*.css'
    ], ['build:dev']).on('change', reload);

});
