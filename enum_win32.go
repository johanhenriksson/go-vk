//go:build windows
// Code generated by go-vk from vk.xml at 2023-02-03 15:01:44.2547785 -0600 CST m=+2.264618601. DO NOT EDIT.
package vk

//go:generate stringer -output=enum_win32_string_0.go -type=FullScreenExclusiveEXT
// FullScreenExclusiveEXT: See https://www.khronos.org/registry/vulkan/specs/1.3-extensions/man/html/VkFullScreenExclusiveEXT.html
type FullScreenExclusiveEXT int32

const (
	FULL_SCREEN_EXCLUSIVE_DEFAULT_EXT                FullScreenExclusiveEXT = 0
	FULL_SCREEN_EXCLUSIVE_ALLOWED_EXT                FullScreenExclusiveEXT = 1
	FULL_SCREEN_EXCLUSIVE_DISALLOWED_EXT             FullScreenExclusiveEXT = 2
	FULL_SCREEN_EXCLUSIVE_APPLICATION_CONTROLLED_EXT FullScreenExclusiveEXT = 3
)

// Platform-specific values for VkResult
const (
	ERROR_FULL_SCREEN_EXCLUSIVE_MODE_LOST_EXT Result = -1000255000
)

// Platform-specific values for VkStructureType
const (
	STRUCTURE_TYPE_FENCE_GET_WIN32_HANDLE_INFO_KHR                StructureType = 1000114002
	STRUCTURE_TYPE_SEMAPHORE_GET_WIN32_HANDLE_INFO_KHR            StructureType = 1000078003
	STRUCTURE_TYPE_IMPORT_FENCE_WIN32_HANDLE_INFO_KHR             StructureType = 1000114000
	STRUCTURE_TYPE_WIN32_KEYED_MUTEX_ACQUIRE_RELEASE_INFO_KHR     StructureType = 1000075000
	STRUCTURE_TYPE_EXPORT_MEMORY_WIN32_HANDLE_INFO_KHR            StructureType = 1000073001
	STRUCTURE_TYPE_IMPORT_MEMORY_WIN32_HANDLE_INFO_KHR            StructureType = 1000073000
	STRUCTURE_TYPE_SURFACE_FULL_SCREEN_EXCLUSIVE_WIN32_INFO_EXT   StructureType = 1000255001
	STRUCTURE_TYPE_MEMORY_WIN32_HANDLE_PROPERTIES_KHR             StructureType = 1000073002
	STRUCTURE_TYPE_MEMORY_GET_WIN32_HANDLE_INFO_KHR               StructureType = 1000073003
	STRUCTURE_TYPE_IMPORT_SEMAPHORE_WIN32_HANDLE_INFO_KHR         StructureType = 1000078000
	STRUCTURE_TYPE_WIN32_SURFACE_CREATE_INFO_KHR                  StructureType = 1000009000
	STRUCTURE_TYPE_SURFACE_CAPABILITIES_FULL_SCREEN_EXCLUSIVE_EXT StructureType = 1000255002
	STRUCTURE_TYPE_EXPORT_MEMORY_WIN32_HANDLE_INFO_NV             StructureType = 1000057001
	STRUCTURE_TYPE_IMPORT_MEMORY_WIN32_HANDLE_INFO_NV             StructureType = 1000057000
	STRUCTURE_TYPE_EXPORT_FENCE_WIN32_HANDLE_INFO_KHR             StructureType = 1000114001
	STRUCTURE_TYPE_WIN32_KEYED_MUTEX_ACQUIRE_RELEASE_INFO_NV      StructureType = 1000058000
	STRUCTURE_TYPE_EXPORT_SEMAPHORE_WIN32_HANDLE_INFO_KHR         StructureType = 1000078001
	STRUCTURE_TYPE_D3D12_FENCE_SUBMIT_INFO_KHR                    StructureType = 1000078002
	STRUCTURE_TYPE_SURFACE_FULL_SCREEN_EXCLUSIVE_INFO_EXT         StructureType = 1000255000
)
