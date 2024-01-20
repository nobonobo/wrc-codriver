set OPENCV_RUNTIME=C:\opencv\build\install\x64\mingw\bin
mkdir dist
del /F /S dist\mark
XCOPY /E /Y voicevox_core dist\voicevox_core\
XCOPY /E /Y assets dist\assets\
XCOPY /E /Y pacenotes dist\pacenotes\
XCOPY /E /Y log dist\log\
COPY /B /Y %OPENCV_RUNTIME%\* dist\
COPY /B /Y onnxruntime.dll dist\onnxruntime.dll
COPY /B /Y wrc-logger.exe dist\wrc-logger.exe
COPY /B /Y base.json dist\base.json
powershell Compress-Archive -Path dist -Force -DestinationPath %1
