import 'dart:ffi';
import 'package:ffi/ffi.dart';
import 'package:ffi/src/utf8.dart';

typedef recognize_func = Pointer<Utf8> Function(Pointer<Utf8> path); // FFI fn signature
typedef Recognize = Pointer<Utf8> Function(Pointer<Utf8> path); // Dart fn signature
final dylib = DynamicLibrary.open('qrcode.a');

final Recognize recognize = dylib.lookup<NativeFunction<recognize_func>>('Recognize').asFunction();

void testffi(String path) {
  print("Hi from dart");
  var value = recognize(path.toNativeUtf8());
  print(value.toDartString());
}

void main(List<String> arguments) {
  //testffi('testdata/image.png');
  testffi('testdata/image2.jpg');
}
