apple:
	@gomobile bind -o ./build/Chisel.xcframework -target=ios,macos,iossimulator -ldflags="-s -w" -v ./
ios:
	@gomobile bind -o ./build/ios/Chisel.xcframework -target=ios,iossimulator -ldflags="-s -w" -v ./
macos:
	@gomobile bind -o ./build/macos/Chisel.xcframework -target=macos -ldflags="-s -w" -v ./