// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/vmware-tanzu/tanzu-cli/pkg/carvelhelpers"
)

type ImageOperationsImpl struct {
	CopyImageFromTarStub        func(string, string) error
	copyImageFromTarMutex       sync.RWMutex
	copyImageFromTarArgsForCall []struct {
		arg1 string
		arg2 string
	}
	copyImageFromTarReturns struct {
		result1 error
	}
	copyImageFromTarReturnsOnCall map[int]struct {
		result1 error
	}
	CopyImageToTarStub        func(string, string) error
	copyImageToTarMutex       sync.RWMutex
	copyImageToTarArgsForCall []struct {
		arg1 string
		arg2 string
	}
	copyImageToTarReturns struct {
		result1 error
	}
	copyImageToTarReturnsOnCall map[int]struct {
		result1 error
	}
	DownloadImageAndSaveFilesToDirStub        func(string, string) error
	downloadImageAndSaveFilesToDirMutex       sync.RWMutex
	downloadImageAndSaveFilesToDirArgsForCall []struct {
		arg1 string
		arg2 string
	}
	downloadImageAndSaveFilesToDirReturns struct {
		result1 error
	}
	downloadImageAndSaveFilesToDirReturnsOnCall map[int]struct {
		result1 error
	}
	GetFileDigestFromImageStub        func(string, string) (string, error)
	getFileDigestFromImageMutex       sync.RWMutex
	getFileDigestFromImageArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getFileDigestFromImageReturns struct {
		result1 string
		result2 error
	}
	getFileDigestFromImageReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	GetFilesMapFromImageStub        func(string) (map[string][]byte, error)
	getFilesMapFromImageMutex       sync.RWMutex
	getFilesMapFromImageArgsForCall []struct {
		arg1 string
	}
	getFilesMapFromImageReturns struct {
		result1 map[string][]byte
		result2 error
	}
	getFilesMapFromImageReturnsOnCall map[int]struct {
		result1 map[string][]byte
		result2 error
	}
	GetImageDigestStub        func(string) (string, string, error)
	getImageDigestMutex       sync.RWMutex
	getImageDigestArgsForCall []struct {
		arg1 string
	}
	getImageDigestReturns struct {
		result1 string
		result2 string
		result3 error
	}
	getImageDigestReturnsOnCall map[int]struct {
		result1 string
		result2 string
		result3 error
	}
	PushImageStub        func(string, []string) error
	pushImageMutex       sync.RWMutex
	pushImageArgsForCall []struct {
		arg1 string
		arg2 []string
	}
	pushImageReturns struct {
		result1 error
	}
	pushImageReturnsOnCall map[int]struct {
		result1 error
	}
	ResolveImageStub        func(string) error
	resolveImageMutex       sync.RWMutex
	resolveImageArgsForCall []struct {
		arg1 string
	}
	resolveImageReturns struct {
		result1 error
	}
	resolveImageReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ImageOperationsImpl) CopyImageFromTar(arg1 string, arg2 string) error {
	fake.copyImageFromTarMutex.Lock()
	ret, specificReturn := fake.copyImageFromTarReturnsOnCall[len(fake.copyImageFromTarArgsForCall)]
	fake.copyImageFromTarArgsForCall = append(fake.copyImageFromTarArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.CopyImageFromTarStub
	fakeReturns := fake.copyImageFromTarReturns
	fake.recordInvocation("CopyImageFromTar", []interface{}{arg1, arg2})
	fake.copyImageFromTarMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *ImageOperationsImpl) CopyImageFromTarCallCount() int {
	fake.copyImageFromTarMutex.RLock()
	defer fake.copyImageFromTarMutex.RUnlock()
	return len(fake.copyImageFromTarArgsForCall)
}

func (fake *ImageOperationsImpl) CopyImageFromTarCalls(stub func(string, string) error) {
	fake.copyImageFromTarMutex.Lock()
	defer fake.copyImageFromTarMutex.Unlock()
	fake.CopyImageFromTarStub = stub
}

func (fake *ImageOperationsImpl) CopyImageFromTarArgsForCall(i int) (string, string) {
	fake.copyImageFromTarMutex.RLock()
	defer fake.copyImageFromTarMutex.RUnlock()
	argsForCall := fake.copyImageFromTarArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *ImageOperationsImpl) CopyImageFromTarReturns(result1 error) {
	fake.copyImageFromTarMutex.Lock()
	defer fake.copyImageFromTarMutex.Unlock()
	fake.CopyImageFromTarStub = nil
	fake.copyImageFromTarReturns = struct {
		result1 error
	}{result1}
}

func (fake *ImageOperationsImpl) CopyImageFromTarReturnsOnCall(i int, result1 error) {
	fake.copyImageFromTarMutex.Lock()
	defer fake.copyImageFromTarMutex.Unlock()
	fake.CopyImageFromTarStub = nil
	if fake.copyImageFromTarReturnsOnCall == nil {
		fake.copyImageFromTarReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.copyImageFromTarReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ImageOperationsImpl) CopyImageToTar(arg1 string, arg2 string) error {
	fake.copyImageToTarMutex.Lock()
	ret, specificReturn := fake.copyImageToTarReturnsOnCall[len(fake.copyImageToTarArgsForCall)]
	fake.copyImageToTarArgsForCall = append(fake.copyImageToTarArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.CopyImageToTarStub
	fakeReturns := fake.copyImageToTarReturns
	fake.recordInvocation("CopyImageToTar", []interface{}{arg1, arg2})
	fake.copyImageToTarMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *ImageOperationsImpl) CopyImageToTarCallCount() int {
	fake.copyImageToTarMutex.RLock()
	defer fake.copyImageToTarMutex.RUnlock()
	return len(fake.copyImageToTarArgsForCall)
}

func (fake *ImageOperationsImpl) CopyImageToTarCalls(stub func(string, string) error) {
	fake.copyImageToTarMutex.Lock()
	defer fake.copyImageToTarMutex.Unlock()
	fake.CopyImageToTarStub = stub
}

func (fake *ImageOperationsImpl) CopyImageToTarArgsForCall(i int) (string, string) {
	fake.copyImageToTarMutex.RLock()
	defer fake.copyImageToTarMutex.RUnlock()
	argsForCall := fake.copyImageToTarArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *ImageOperationsImpl) CopyImageToTarReturns(result1 error) {
	fake.copyImageToTarMutex.Lock()
	defer fake.copyImageToTarMutex.Unlock()
	fake.CopyImageToTarStub = nil
	fake.copyImageToTarReturns = struct {
		result1 error
	}{result1}
}

func (fake *ImageOperationsImpl) CopyImageToTarReturnsOnCall(i int, result1 error) {
	fake.copyImageToTarMutex.Lock()
	defer fake.copyImageToTarMutex.Unlock()
	fake.CopyImageToTarStub = nil
	if fake.copyImageToTarReturnsOnCall == nil {
		fake.copyImageToTarReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.copyImageToTarReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ImageOperationsImpl) DownloadImageAndSaveFilesToDir(arg1 string, arg2 string) error {
	fake.downloadImageAndSaveFilesToDirMutex.Lock()
	ret, specificReturn := fake.downloadImageAndSaveFilesToDirReturnsOnCall[len(fake.downloadImageAndSaveFilesToDirArgsForCall)]
	fake.downloadImageAndSaveFilesToDirArgsForCall = append(fake.downloadImageAndSaveFilesToDirArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.DownloadImageAndSaveFilesToDirStub
	fakeReturns := fake.downloadImageAndSaveFilesToDirReturns
	fake.recordInvocation("DownloadImageAndSaveFilesToDir", []interface{}{arg1, arg2})
	fake.downloadImageAndSaveFilesToDirMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *ImageOperationsImpl) DownloadImageAndSaveFilesToDirCallCount() int {
	fake.downloadImageAndSaveFilesToDirMutex.RLock()
	defer fake.downloadImageAndSaveFilesToDirMutex.RUnlock()
	return len(fake.downloadImageAndSaveFilesToDirArgsForCall)
}

func (fake *ImageOperationsImpl) DownloadImageAndSaveFilesToDirCalls(stub func(string, string) error) {
	fake.downloadImageAndSaveFilesToDirMutex.Lock()
	defer fake.downloadImageAndSaveFilesToDirMutex.Unlock()
	fake.DownloadImageAndSaveFilesToDirStub = stub
}

func (fake *ImageOperationsImpl) DownloadImageAndSaveFilesToDirArgsForCall(i int) (string, string) {
	fake.downloadImageAndSaveFilesToDirMutex.RLock()
	defer fake.downloadImageAndSaveFilesToDirMutex.RUnlock()
	argsForCall := fake.downloadImageAndSaveFilesToDirArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *ImageOperationsImpl) DownloadImageAndSaveFilesToDirReturns(result1 error) {
	fake.downloadImageAndSaveFilesToDirMutex.Lock()
	defer fake.downloadImageAndSaveFilesToDirMutex.Unlock()
	fake.DownloadImageAndSaveFilesToDirStub = nil
	fake.downloadImageAndSaveFilesToDirReturns = struct {
		result1 error
	}{result1}
}

func (fake *ImageOperationsImpl) DownloadImageAndSaveFilesToDirReturnsOnCall(i int, result1 error) {
	fake.downloadImageAndSaveFilesToDirMutex.Lock()
	defer fake.downloadImageAndSaveFilesToDirMutex.Unlock()
	fake.DownloadImageAndSaveFilesToDirStub = nil
	if fake.downloadImageAndSaveFilesToDirReturnsOnCall == nil {
		fake.downloadImageAndSaveFilesToDirReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.downloadImageAndSaveFilesToDirReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ImageOperationsImpl) GetFileDigestFromImage(arg1 string, arg2 string) (string, error) {
	fake.getFileDigestFromImageMutex.Lock()
	ret, specificReturn := fake.getFileDigestFromImageReturnsOnCall[len(fake.getFileDigestFromImageArgsForCall)]
	fake.getFileDigestFromImageArgsForCall = append(fake.getFileDigestFromImageArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.GetFileDigestFromImageStub
	fakeReturns := fake.getFileDigestFromImageReturns
	fake.recordInvocation("GetFileDigestFromImage", []interface{}{arg1, arg2})
	fake.getFileDigestFromImageMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ImageOperationsImpl) GetFileDigestFromImageCallCount() int {
	fake.getFileDigestFromImageMutex.RLock()
	defer fake.getFileDigestFromImageMutex.RUnlock()
	return len(fake.getFileDigestFromImageArgsForCall)
}

func (fake *ImageOperationsImpl) GetFileDigestFromImageCalls(stub func(string, string) (string, error)) {
	fake.getFileDigestFromImageMutex.Lock()
	defer fake.getFileDigestFromImageMutex.Unlock()
	fake.GetFileDigestFromImageStub = stub
}

func (fake *ImageOperationsImpl) GetFileDigestFromImageArgsForCall(i int) (string, string) {
	fake.getFileDigestFromImageMutex.RLock()
	defer fake.getFileDigestFromImageMutex.RUnlock()
	argsForCall := fake.getFileDigestFromImageArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *ImageOperationsImpl) GetFileDigestFromImageReturns(result1 string, result2 error) {
	fake.getFileDigestFromImageMutex.Lock()
	defer fake.getFileDigestFromImageMutex.Unlock()
	fake.GetFileDigestFromImageStub = nil
	fake.getFileDigestFromImageReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *ImageOperationsImpl) GetFileDigestFromImageReturnsOnCall(i int, result1 string, result2 error) {
	fake.getFileDigestFromImageMutex.Lock()
	defer fake.getFileDigestFromImageMutex.Unlock()
	fake.GetFileDigestFromImageStub = nil
	if fake.getFileDigestFromImageReturnsOnCall == nil {
		fake.getFileDigestFromImageReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getFileDigestFromImageReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *ImageOperationsImpl) GetFilesMapFromImage(arg1 string) (map[string][]byte, error) {
	fake.getFilesMapFromImageMutex.Lock()
	ret, specificReturn := fake.getFilesMapFromImageReturnsOnCall[len(fake.getFilesMapFromImageArgsForCall)]
	fake.getFilesMapFromImageArgsForCall = append(fake.getFilesMapFromImageArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetFilesMapFromImageStub
	fakeReturns := fake.getFilesMapFromImageReturns
	fake.recordInvocation("GetFilesMapFromImage", []interface{}{arg1})
	fake.getFilesMapFromImageMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ImageOperationsImpl) GetFilesMapFromImageCallCount() int {
	fake.getFilesMapFromImageMutex.RLock()
	defer fake.getFilesMapFromImageMutex.RUnlock()
	return len(fake.getFilesMapFromImageArgsForCall)
}

func (fake *ImageOperationsImpl) GetFilesMapFromImageCalls(stub func(string) (map[string][]byte, error)) {
	fake.getFilesMapFromImageMutex.Lock()
	defer fake.getFilesMapFromImageMutex.Unlock()
	fake.GetFilesMapFromImageStub = stub
}

func (fake *ImageOperationsImpl) GetFilesMapFromImageArgsForCall(i int) string {
	fake.getFilesMapFromImageMutex.RLock()
	defer fake.getFilesMapFromImageMutex.RUnlock()
	argsForCall := fake.getFilesMapFromImageArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ImageOperationsImpl) GetFilesMapFromImageReturns(result1 map[string][]byte, result2 error) {
	fake.getFilesMapFromImageMutex.Lock()
	defer fake.getFilesMapFromImageMutex.Unlock()
	fake.GetFilesMapFromImageStub = nil
	fake.getFilesMapFromImageReturns = struct {
		result1 map[string][]byte
		result2 error
	}{result1, result2}
}

func (fake *ImageOperationsImpl) GetFilesMapFromImageReturnsOnCall(i int, result1 map[string][]byte, result2 error) {
	fake.getFilesMapFromImageMutex.Lock()
	defer fake.getFilesMapFromImageMutex.Unlock()
	fake.GetFilesMapFromImageStub = nil
	if fake.getFilesMapFromImageReturnsOnCall == nil {
		fake.getFilesMapFromImageReturnsOnCall = make(map[int]struct {
			result1 map[string][]byte
			result2 error
		})
	}
	fake.getFilesMapFromImageReturnsOnCall[i] = struct {
		result1 map[string][]byte
		result2 error
	}{result1, result2}
}

func (fake *ImageOperationsImpl) GetImageDigest(arg1 string) (string, string, error) {
	fake.getImageDigestMutex.Lock()
	ret, specificReturn := fake.getImageDigestReturnsOnCall[len(fake.getImageDigestArgsForCall)]
	fake.getImageDigestArgsForCall = append(fake.getImageDigestArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetImageDigestStub
	fakeReturns := fake.getImageDigestReturns
	fake.recordInvocation("GetImageDigest", []interface{}{arg1})
	fake.getImageDigestMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *ImageOperationsImpl) GetImageDigestCallCount() int {
	fake.getImageDigestMutex.RLock()
	defer fake.getImageDigestMutex.RUnlock()
	return len(fake.getImageDigestArgsForCall)
}

func (fake *ImageOperationsImpl) GetImageDigestCalls(stub func(string) (string, string, error)) {
	fake.getImageDigestMutex.Lock()
	defer fake.getImageDigestMutex.Unlock()
	fake.GetImageDigestStub = stub
}

func (fake *ImageOperationsImpl) GetImageDigestArgsForCall(i int) string {
	fake.getImageDigestMutex.RLock()
	defer fake.getImageDigestMutex.RUnlock()
	argsForCall := fake.getImageDigestArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ImageOperationsImpl) GetImageDigestReturns(result1 string, result2 string, result3 error) {
	fake.getImageDigestMutex.Lock()
	defer fake.getImageDigestMutex.Unlock()
	fake.GetImageDigestStub = nil
	fake.getImageDigestReturns = struct {
		result1 string
		result2 string
		result3 error
	}{result1, result2, result3}
}

func (fake *ImageOperationsImpl) GetImageDigestReturnsOnCall(i int, result1 string, result2 string, result3 error) {
	fake.getImageDigestMutex.Lock()
	defer fake.getImageDigestMutex.Unlock()
	fake.GetImageDigestStub = nil
	if fake.getImageDigestReturnsOnCall == nil {
		fake.getImageDigestReturnsOnCall = make(map[int]struct {
			result1 string
			result2 string
			result3 error
		})
	}
	fake.getImageDigestReturnsOnCall[i] = struct {
		result1 string
		result2 string
		result3 error
	}{result1, result2, result3}
}

func (fake *ImageOperationsImpl) PushImage(arg1 string, arg2 []string) error {
	var arg2Copy []string
	if arg2 != nil {
		arg2Copy = make([]string, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.pushImageMutex.Lock()
	ret, specificReturn := fake.pushImageReturnsOnCall[len(fake.pushImageArgsForCall)]
	fake.pushImageArgsForCall = append(fake.pushImageArgsForCall, struct {
		arg1 string
		arg2 []string
	}{arg1, arg2Copy})
	stub := fake.PushImageStub
	fakeReturns := fake.pushImageReturns
	fake.recordInvocation("PushImage", []interface{}{arg1, arg2Copy})
	fake.pushImageMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *ImageOperationsImpl) PushImageCallCount() int {
	fake.pushImageMutex.RLock()
	defer fake.pushImageMutex.RUnlock()
	return len(fake.pushImageArgsForCall)
}

func (fake *ImageOperationsImpl) PushImageCalls(stub func(string, []string) error) {
	fake.pushImageMutex.Lock()
	defer fake.pushImageMutex.Unlock()
	fake.PushImageStub = stub
}

func (fake *ImageOperationsImpl) PushImageArgsForCall(i int) (string, []string) {
	fake.pushImageMutex.RLock()
	defer fake.pushImageMutex.RUnlock()
	argsForCall := fake.pushImageArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *ImageOperationsImpl) PushImageReturns(result1 error) {
	fake.pushImageMutex.Lock()
	defer fake.pushImageMutex.Unlock()
	fake.PushImageStub = nil
	fake.pushImageReturns = struct {
		result1 error
	}{result1}
}

func (fake *ImageOperationsImpl) PushImageReturnsOnCall(i int, result1 error) {
	fake.pushImageMutex.Lock()
	defer fake.pushImageMutex.Unlock()
	fake.PushImageStub = nil
	if fake.pushImageReturnsOnCall == nil {
		fake.pushImageReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.pushImageReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ImageOperationsImpl) ResolveImage(arg1 string) error {
	fake.resolveImageMutex.Lock()
	ret, specificReturn := fake.resolveImageReturnsOnCall[len(fake.resolveImageArgsForCall)]
	fake.resolveImageArgsForCall = append(fake.resolveImageArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.ResolveImageStub
	fakeReturns := fake.resolveImageReturns
	fake.recordInvocation("ResolveImage", []interface{}{arg1})
	fake.resolveImageMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *ImageOperationsImpl) ResolveImageCallCount() int {
	fake.resolveImageMutex.RLock()
	defer fake.resolveImageMutex.RUnlock()
	return len(fake.resolveImageArgsForCall)
}

func (fake *ImageOperationsImpl) ResolveImageCalls(stub func(string) error) {
	fake.resolveImageMutex.Lock()
	defer fake.resolveImageMutex.Unlock()
	fake.ResolveImageStub = stub
}

func (fake *ImageOperationsImpl) ResolveImageArgsForCall(i int) string {
	fake.resolveImageMutex.RLock()
	defer fake.resolveImageMutex.RUnlock()
	argsForCall := fake.resolveImageArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ImageOperationsImpl) ResolveImageReturns(result1 error) {
	fake.resolveImageMutex.Lock()
	defer fake.resolveImageMutex.Unlock()
	fake.ResolveImageStub = nil
	fake.resolveImageReturns = struct {
		result1 error
	}{result1}
}

func (fake *ImageOperationsImpl) ResolveImageReturnsOnCall(i int, result1 error) {
	fake.resolveImageMutex.Lock()
	defer fake.resolveImageMutex.Unlock()
	fake.ResolveImageStub = nil
	if fake.resolveImageReturnsOnCall == nil {
		fake.resolveImageReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.resolveImageReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ImageOperationsImpl) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.copyImageFromTarMutex.RLock()
	defer fake.copyImageFromTarMutex.RUnlock()
	fake.copyImageToTarMutex.RLock()
	defer fake.copyImageToTarMutex.RUnlock()
	fake.downloadImageAndSaveFilesToDirMutex.RLock()
	defer fake.downloadImageAndSaveFilesToDirMutex.RUnlock()
	fake.getFileDigestFromImageMutex.RLock()
	defer fake.getFileDigestFromImageMutex.RUnlock()
	fake.getFilesMapFromImageMutex.RLock()
	defer fake.getFilesMapFromImageMutex.RUnlock()
	fake.getImageDigestMutex.RLock()
	defer fake.getImageDigestMutex.RUnlock()
	fake.pushImageMutex.RLock()
	defer fake.pushImageMutex.RUnlock()
	fake.resolveImageMutex.RLock()
	defer fake.resolveImageMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *ImageOperationsImpl) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ carvelhelpers.ImageOperationsImpl = new(ImageOperationsImpl)
