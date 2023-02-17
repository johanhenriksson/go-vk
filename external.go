// Code generated by go-vk from vk.xml at 2023-02-03 15:01:43.3815213 -0600 CST m=+1.391361401. DO NOT EDIT.
package vk

const (
	LOD_CLAMP_NONE float32 = 1000.0
)

const (
	FALSE     uint32 = 0
	TRUE      uint32 = 1
	LUID_SIZE uint32 = 8
	// The maximum number of unique memory heaps, each of which supporting 1 or more memory types
	MAX_MEMORY_HEAPS              uint32 = 16
	UUID_SIZE                     uint32 = 16
	MAX_GLOBAL_PRIORITY_SIZE_EXT  uint32 = 16
	MAX_MEMORY_TYPES              uint32 = 32
	MAX_DEVICE_GROUP_SIZE         uint32 = 32
	MAX_EXTENSION_NAME_SIZE       uint32 = 256
	MAX_DRIVER_INFO_SIZE          uint32 = 256
	MAX_PHYSICAL_DEVICE_NAME_SIZE uint32 = 256
	MAX_DESCRIPTION_SIZE          uint32 = 256
	MAX_DRIVER_NAME_SIZE          uint32 = 256
	ATTACHMENT_UNUSED             uint32 = ^uint32(0)
	SUBPASS_EXTERNAL              uint32 = ^uint32(0)
	REMAINING_ARRAY_LAYERS        uint32 = ^uint32(0)
	QUEUE_FAMILY_IGNORED          uint32 = ^uint32(0)
	REMAINING_MIP_LEVELS          uint32 = ^uint32(0)
	SHADER_UNUSED_KHR             uint32 = ^uint32(0)
	QUEUE_FAMILY_EXTERNAL         uint32 = ^uint32(1)
	QUEUE_FAMILY_FOREIGN_EXT      uint32 = ^uint32(2)
)

const (
	WHOLE_SIZE uint64 = ^uint64(0)
)

// Platform-specific values for
const (
	KHR_DEFERRED_HOST_OPERATIONS_SPEC_VERSION                      = 4
	EXT_BUFFER_DEVICE_ADDRESS_SPEC_VERSION                         = 2
	NV_RAY_TRACING_MOTION_BLUR_SPEC_VERSION                        = 1
	EXT_LOAD_STORE_OP_NONE_SPEC_VERSION                            = 1
	EXT_PCI_BUS_INFO_SPEC_VERSION                                  = 2
	KHR_DISPLAY_SWAPCHAIN_EXTENSION_NAME                           = "VK_KHR_display_swapchain"
	KHR_DEDICATED_ALLOCATION_SPEC_VERSION                          = 3
	KHR_RELAXED_BLOCK_LAYOUT_SPEC_VERSION                          = 1
	EXT_ACQUIRE_DRM_DISPLAY_SPEC_VERSION                           = 1
	QCOM_RENDER_PASS_TRANSFORM_EXTENSION_NAME                      = "VK_QCOM_render_pass_transform"
	KHR_INCREMENTAL_PRESENT_EXTENSION_NAME                         = "VK_KHR_incremental_present"
	EXT_INLINE_UNIFORM_BLOCK_SPEC_VERSION                          = 1
	KHR_RAY_TRACING_PIPELINE_EXTENSION_NAME                        = "VK_KHR_ray_tracing_pipeline"
	KHR_PRESENT_WAIT_SPEC_VERSION                                  = 1
	EXT_ROBUSTNESS_2_EXTENSION_NAME                                = "VK_EXT_robustness2"
	AMD_MIXED_ATTACHMENT_SAMPLES_EXTENSION_NAME                    = "VK_AMD_mixed_attachment_samples"
	KHR_STORAGE_BUFFER_STORAGE_CLASS_SPEC_VERSION                  = 1
	KHR_ACCELERATION_STRUCTURE_SPEC_VERSION                        = 12
	KHR_DEVICE_GROUP_CREATION_EXTENSION_NAME                       = "VK_KHR_device_group_creation"
	EXT_MEMORY_BUDGET_SPEC_VERSION                                 = 1
	EXT_SHADER_ATOMIC_FLOAT_EXTENSION_NAME                         = "VK_EXT_shader_atomic_float"
	KHR_BUFFER_DEVICE_ADDRESS_EXTENSION_NAME                       = "VK_KHR_buffer_device_address"
	KHR_WORKGROUP_MEMORY_EXPLICIT_LAYOUT_SPEC_VERSION              = 1
	EXT_HOST_QUERY_RESET_EXTENSION_NAME                            = "VK_EXT_host_query_reset"
	KHR_MAINTENANCE3_SPEC_VERSION                                  = 1
	NV_FILL_RECTANGLE_SPEC_VERSION                                 = 1
	EXT_FRAGMENT_SHADER_INTERLOCK_SPEC_VERSION                     = 1
	AMD_SHADER_EXPLICIT_VERTEX_PARAMETER_SPEC_VERSION              = 1
	KHR_SHADER_SUBGROUP_UNIFORM_CONTROL_FLOW_EXTENSION_NAME        = "VK_KHR_shader_subgroup_uniform_control_flow"
	KHR_EXTERNAL_FENCE_SPEC_VERSION                                = 1
	NV_FRAGMENT_COVERAGE_TO_COLOR_EXTENSION_NAME                   = "VK_NV_fragment_coverage_to_color"
	EXT_CUSTOM_BORDER_COLOR_EXTENSION_NAME                         = "VK_EXT_custom_border_color"
	NV_MESH_SHADER_EXTENSION_NAME                                  = "VK_NV_mesh_shader"
	EXT_TOOLING_INFO_SPEC_VERSION                                  = 1
	NV_FRAMEBUFFER_MIXED_SAMPLES_SPEC_VERSION                      = 1
	EXT_HDR_METADATA_SPEC_VERSION                                  = 2
	AMD_SHADER_TRINARY_MINMAX_SPEC_VERSION                         = 1
	EXT_BUFFER_DEVICE_ADDRESS_EXTENSION_NAME                       = "VK_EXT_buffer_device_address"
	EXT_PROVOKING_VERTEX_SPEC_VERSION                              = 1
	EXT_PIPELINE_CREATION_CACHE_CONTROL_SPEC_VERSION               = 3
	EXT_DEPTH_RANGE_UNRESTRICTED_EXTENSION_NAME                    = "VK_EXT_depth_range_unrestricted"
	NV_CORNER_SAMPLED_IMAGE_SPEC_VERSION                           = 2
	KHR_UNIFORM_BUFFER_STANDARD_LAYOUT_EXTENSION_NAME              = "VK_KHR_uniform_buffer_standard_layout"
	EXT_GLOBAL_PRIORITY_QUERY_EXTENSION_NAME                       = "VK_EXT_global_priority_query"
	EXT_SHADER_IMAGE_ATOMIC_INT64_SPEC_VERSION                     = 1
	NV_VIEWPORT_SWIZZLE_EXTENSION_NAME                             = "VK_NV_viewport_swizzle"
	KHR_SHADER_DRAW_PARAMETERS_SPEC_VERSION                        = 1
	NV_EXTERNAL_MEMORY_SPEC_VERSION                                = 1
	AMD_MIXED_ATTACHMENT_SAMPLES_SPEC_VERSION                      = 1
	KHR_EXTERNAL_FENCE_FD_SPEC_VERSION                             = 1
	KHR_GET_PHYSICAL_DEVICE_PROPERTIES_2_EXTENSION_NAME            = "VK_KHR_get_physical_device_properties2"
	EXT_SHADER_STENCIL_EXPORT_SPEC_VERSION                         = 1
	NV_GEOMETRY_SHADER_PASSTHROUGH_SPEC_VERSION                    = 1
	KHR_PRESENT_WAIT_EXTENSION_NAME                                = "VK_KHR_present_wait"
	NV_CLIP_SPACE_W_SCALING_EXTENSION_NAME                         = "VK_NV_clip_space_w_scaling"
	NV_CORNER_SAMPLED_IMAGE_EXTENSION_NAME                         = "VK_NV_corner_sampled_image"
	AMD_DISPLAY_NATIVE_HDR_EXTENSION_NAME                          = "VK_AMD_display_native_hdr"
	KHR_SHADER_SUBGROUP_UNIFORM_CONTROL_FLOW_SPEC_VERSION          = 1
	EXT_SHADER_ATOMIC_FLOAT_2_EXTENSION_NAME                       = "VK_EXT_shader_atomic_float2"
	NV_GLSL_SHADER_EXTENSION_NAME                                  = "VK_NV_glsl_shader"
	KHR_SAMPLER_YCBCR_CONVERSION_EXTENSION_NAME                    = "VK_KHR_sampler_ycbcr_conversion"
	NVX_BINARY_IMPORT_EXTENSION_NAME                               = "VK_NVX_binary_import"
	SHADER_UNUSED_NV                                        uint32 = SHADER_UNUSED_KHR
	NV_EXTERNAL_MEMORY_CAPABILITIES_SPEC_VERSION                   = 1
	KHR_DISPLAY_EXTENSION_NAME                                     = "VK_KHR_display"
	KHR_RAY_QUERY_EXTENSION_NAME                                   = "VK_KHR_ray_query"
	AMD_SHADER_CORE_PROPERTIES_SPEC_VERSION                        = 2
	EXT_GLOBAL_PRIORITY_SPEC_VERSION                               = 2
	KHR_SHARED_PRESENTABLE_IMAGE_SPEC_VERSION                      = 1
	NV_SHADER_IMAGE_FOOTPRINT_SPEC_VERSION                         = 2
	EXT_SAMPLER_FILTER_MINMAX_SPEC_VERSION                         = 2
	KHR_GET_DISPLAY_PROPERTIES_2_EXTENSION_NAME                    = "VK_KHR_get_display_properties2"
	NV_FRAGMENT_COVERAGE_TO_COLOR_SPEC_VERSION                     = 1
	EXT_DEPTH_RANGE_UNRESTRICTED_SPEC_VERSION                      = 1
	NV_DEVICE_GENERATED_COMMANDS_SPEC_VERSION                      = 3
	EXT_MULTI_DRAW_EXTENSION_NAME                                  = "VK_EXT_multi_draw"
	EXT_ACQUIRE_DRM_DISPLAY_EXTENSION_NAME                         = "VK_EXT_acquire_drm_display"
	EXT_FRAGMENT_DENSITY_MAP_2_SPEC_VERSION                        = 1
	KHR_TIMELINE_SEMAPHORE_SPEC_VERSION                            = 2
	EXT_CONDITIONAL_RENDERING_SPEC_VERSION                         = 2
	GOOGLE_USER_TYPE_SPEC_VERSION                                  = 1
	EXT_DEPTH_CLIP_ENABLE_EXTENSION_NAME                           = "VK_EXT_depth_clip_enable"
	EXT_SHADER_DEMOTE_TO_HELPER_INVOCATION_SPEC_VERSION            = 1
	IMG_FILTER_CUBIC_SPEC_VERSION                                  = 1
	EXT_VALIDATION_CACHE_EXTENSION_NAME                            = "VK_EXT_validation_cache"
	QUEUE_FAMILY_EXTERNAL_KHR                               uint32 = QUEUE_FAMILY_EXTERNAL
	EXT_VALIDATION_CACHE_SPEC_VERSION                              = 1
	EXT_IMAGE_ROBUSTNESS_SPEC_VERSION                              = 1
	KHR_TIMELINE_SEMAPHORE_EXTENSION_NAME                          = "VK_KHR_timeline_semaphore"
	KHR_EXTERNAL_SEMAPHORE_EXTENSION_NAME                          = "VK_KHR_external_semaphore"
	EXT_QUEUE_FAMILY_FOREIGN_SPEC_VERSION                          = 1
	EXT_DESCRIPTOR_INDEXING_EXTENSION_NAME                         = "VK_EXT_descriptor_indexing"
	EXT_DEVICE_MEMORY_REPORT_SPEC_VERSION                          = 2
	KHR_PIPELINE_EXECUTABLE_PROPERTIES_SPEC_VERSION                = 1
	MAX_DRIVER_NAME_SIZE_KHR                                uint32 = MAX_DRIVER_NAME_SIZE
	EXT_TEXEL_BUFFER_ALIGNMENT_SPEC_VERSION                        = 1
	KHR_SAMPLER_MIRROR_CLAMP_TO_EDGE_EXTENSION_NAME                = "VK_KHR_sampler_mirror_clamp_to_edge"
	INTEL_PERFORMANCE_QUERY_EXTENSION_NAME                         = "VK_INTEL_performance_query"
	KHR_SYNCHRONIZATION_2_SPEC_VERSION                             = 1
	MAX_DRIVER_INFO_SIZE_KHR                                uint32 = MAX_DRIVER_INFO_SIZE
	EXT_INDEX_TYPE_UINT8_SPEC_VERSION                              = 1
	KHR_EXTERNAL_MEMORY_SPEC_VERSION                               = 1
	EXT_TEXEL_BUFFER_ALIGNMENT_EXTENSION_NAME                      = "VK_EXT_texel_buffer_alignment"
	NV_SHADER_SM_BUILTINS_SPEC_VERSION                             = 1
	NVX_IMAGE_VIEW_HANDLE_SPEC_VERSION                             = 2
	EXT_PRIVATE_DATA_SPEC_VERSION                                  = 1
	GOOGLE_DECORATE_STRING_SPEC_VERSION                            = 1
	QCOM_RENDER_PASS_SHADER_RESOLVE_EXTENSION_NAME                 = "VK_QCOM_render_pass_shader_resolve"
	AMD_SHADER_FRAGMENT_MASK_SPEC_VERSION                          = 1
	EXT_LINE_RASTERIZATION_SPEC_VERSION                            = 1
	EXT_IMAGE_DRM_FORMAT_MODIFIER_SPEC_VERSION                     = 1
	NV_GLSL_SHADER_SPEC_VERSION                                    = 1
	AMD_GPU_SHADER_INT16_EXTENSION_NAME                            = "VK_AMD_gpu_shader_int16"
	KHR_DEFERRED_HOST_OPERATIONS_EXTENSION_NAME                    = "VK_KHR_deferred_host_operations"
	EXT_HEADLESS_SURFACE_SPEC_VERSION                              = 1
	KHR_ZERO_INITIALIZE_WORKGROUP_MEMORY_EXTENSION_NAME            = "VK_KHR_zero_initialize_workgroup_memory"
	NV_VIEWPORT_SWIZZLE_SPEC_VERSION                               = 1
	EXT_SAMPLE_LOCATIONS_EXTENSION_NAME                            = "VK_EXT_sample_locations"
	NV_EXTERNAL_MEMORY_CAPABILITIES_EXTENSION_NAME                 = "VK_NV_external_memory_capabilities"
	EXT_4444_FORMATS_SPEC_VERSION                                  = 1
	INTEL_PERFORMANCE_QUERY_SPEC_VERSION                           = 2
	NV_FILL_RECTANGLE_EXTENSION_NAME                               = "VK_NV_fill_rectangle"
	KHR_DEVICE_GROUP_SPEC_VERSION                                  = 4
	KHR_EXTERNAL_SEMAPHORE_CAPABILITIES_EXTENSION_NAME             = "VK_KHR_external_semaphore_capabilities"
	HUAWEI_INVOCATION_MASK_SPEC_VERSION                            = 1
	EXT_HEADLESS_SURFACE_EXTENSION_NAME                            = "VK_EXT_headless_surface"
	KHR_DESCRIPTOR_UPDATE_TEMPLATE_EXTENSION_NAME                  = "VK_KHR_descriptor_update_template"
	EXT_GLOBAL_PRIORITY_QUERY_SPEC_VERSION                         = 1
	KHR_SHADER_NON_SEMANTIC_INFO_SPEC_VERSION                      = 1
	QCOM_RENDER_PASS_STORE_OPS_EXTENSION_NAME                      = "VK_QCOM_render_pass_store_ops"
	EXT_COLOR_WRITE_ENABLE_EXTENSION_NAME                          = "VK_EXT_color_write_enable"
	MAX_DEVICE_GROUP_SIZE_KHR                               uint32 = MAX_DEVICE_GROUP_SIZE
	KHR_RELAXED_BLOCK_LAYOUT_EXTENSION_NAME                        = "VK_KHR_relaxed_block_layout"
	EXT_FRAGMENT_DENSITY_MAP_SPEC_VERSION                          = 1
	KHR_PERFORMANCE_QUERY_SPEC_VERSION                             = 1
	QCOM_RENDER_PASS_STORE_OPS_SPEC_VERSION                        = 2
	KHR_IMAGE_FORMAT_LIST_EXTENSION_NAME                           = "VK_KHR_image_format_list"
	EXT_CONSERVATIVE_RASTERIZATION_SPEC_VERSION                    = 1
	EXT_MEMORY_BUDGET_EXTENSION_NAME                               = "VK_EXT_memory_budget"
	NV_COOPERATIVE_MATRIX_SPEC_VERSION                             = 1
	EXT_SWAPCHAIN_COLOR_SPACE_SPEC_VERSION                         = 4
	KHR_GET_DISPLAY_PROPERTIES_2_SPEC_VERSION                      = 1
	KHR_VARIABLE_POINTERS_EXTENSION_NAME                           = "VK_KHR_variable_pointers"
	NV_COVERAGE_REDUCTION_MODE_EXTENSION_NAME                      = "VK_NV_coverage_reduction_mode"
	AMD_DRAW_INDIRECT_COUNT_SPEC_VERSION                           = 2
	NV_INHERITED_VIEWPORT_SCISSOR_EXTENSION_NAME                   = "VK_NV_inherited_viewport_scissor"
	KHR_GET_MEMORY_REQUIREMENTS_2_SPEC_VERSION                     = 1
	AMD_GPU_SHADER_INT16_SPEC_VERSION                              = 2
	EXT_BLEND_OPERATION_ADVANCED_EXTENSION_NAME                    = "VK_EXT_blend_operation_advanced"
	EXT_VALIDATION_FEATURES_EXTENSION_NAME                         = "VK_EXT_validation_features"
	EXT_EXTENDED_DYNAMIC_STATE_2_EXTENSION_NAME                    = "VK_EXT_extended_dynamic_state2"
	EXT_PIPELINE_CREATION_FEEDBACK_SPEC_VERSION                    = 1
	VALVE_MUTABLE_DESCRIPTOR_TYPE_EXTENSION_NAME                   = "VK_VALVE_mutable_descriptor_type"
	KHR_SHARED_PRESENTABLE_IMAGE_EXTENSION_NAME                    = "VK_KHR_shared_presentable_image"
	EXT_GLOBAL_PRIORITY_EXTENSION_NAME                             = "VK_EXT_global_priority"
	NV_SHADER_SUBGROUP_PARTITIONED_SPEC_VERSION                    = 1
	EXT_TRANSFORM_FEEDBACK_SPEC_VERSION                            = 1
	EXT_ASTC_DECODE_MODE_SPEC_VERSION                              = 1
	EXT_SWAPCHAIN_COLOR_SPACE_EXTENSION_NAME                       = "VK_EXT_swapchain_colorspace"
	NV_RAY_TRACING_EXTENSION_NAME                                  = "VK_NV_ray_tracing"
	NV_DEVICE_DIAGNOSTIC_CHECKPOINTS_EXTENSION_NAME                = "VK_NV_device_diagnostic_checkpoints"
	EXT_PHYSICAL_DEVICE_DRM_EXTENSION_NAME                         = "VK_EXT_physical_device_drm"
	EXT_FRAGMENT_DENSITY_MAP_EXTENSION_NAME                        = "VK_EXT_fragment_density_map"
	EXT_SHADER_IMAGE_ATOMIC_INT64_EXTENSION_NAME                   = "VK_EXT_shader_image_atomic_int64"
	AMD_GCN_SHADER_EXTENSION_NAME                                  = "VK_AMD_gcn_shader"
	KHR_FRAGMENT_SHADING_RATE_SPEC_VERSION                         = 1
	KHR_SHADER_SUBGROUP_EXTENDED_TYPES_SPEC_VERSION                = 1
	KHR_EXTERNAL_SEMAPHORE_SPEC_VERSION                            = 1
	EXT_DEBUG_MARKER_SPEC_VERSION                                  = 4
	EXT_DEBUG_MARKER_EXTENSION_NAME                                = "VK_EXT_debug_marker"
	KHR_DRIVER_PROPERTIES_EXTENSION_NAME                           = "VK_KHR_driver_properties"
	EXT_EXTERNAL_MEMORY_DMA_BUF_EXTENSION_NAME                     = "VK_EXT_external_memory_dma_buf"
	EXT_VALIDATION_FLAGS_EXTENSION_NAME                            = "VK_EXT_validation_flags"
	KHR_PRESENT_ID_EXTENSION_NAME                                  = "VK_KHR_present_id"
	EXT_SHADER_VIEWPORT_INDEX_LAYER_SPEC_VERSION                   = 1
	NV_DEVICE_DIAGNOSTIC_CHECKPOINTS_SPEC_VERSION                  = 2
	KHR_DRAW_INDIRECT_COUNT_EXTENSION_NAME                         = "VK_KHR_draw_indirect_count"
	EXT_MEMORY_PRIORITY_SPEC_VERSION                               = 1
	EXT_SUBGROUP_SIZE_CONTROL_SPEC_VERSION                         = 2
	KHR_STORAGE_BUFFER_STORAGE_CLASS_EXTENSION_NAME                = "VK_KHR_storage_buffer_storage_class"
	EXT_DISCARD_RECTANGLES_EXTENSION_NAME                          = "VK_EXT_discard_rectangles"
	NV_GEOMETRY_SHADER_PASSTHROUGH_EXTENSION_NAME                  = "VK_NV_geometry_shader_passthrough"
	AMD_NEGATIVE_VIEWPORT_HEIGHT_SPEC_VERSION                      = 1
	KHR_MAINTENANCE2_SPEC_VERSION                                  = 1
	IMG_FILTER_CUBIC_EXTENSION_NAME                                = "VK_IMG_filter_cubic"
	EXT_SHADER_ATOMIC_FLOAT_SPEC_VERSION                           = 1
	HUAWEI_SUBPASS_SHADING_EXTENSION_NAME                          = "VK_HUAWEI_subpass_shading"
	EXT_PROVOKING_VERTEX_EXTENSION_NAME                            = "VK_EXT_provoking_vertex"
	KHR_MULTIVIEW_SPEC_VERSION                                     = 1
	EXT_FILTER_CUBIC_EXTENSION_NAME                                = "VK_EXT_filter_cubic"
	KHR_EXTERNAL_FENCE_FD_EXTENSION_NAME                           = "VK_KHR_external_fence_fd"
	KHR_BIND_MEMORY_2_EXTENSION_NAME                               = "VK_KHR_bind_memory2"
	EXT_YCBCR_2PLANE_444_FORMATS_SPEC_VERSION                      = 1
	NV_COVERAGE_REDUCTION_MODE_SPEC_VERSION                        = 1
	AMD_SHADER_IMAGE_LOAD_STORE_LOD_EXTENSION_NAME                 = "VK_AMD_shader_image_load_store_lod"
	AMD_MEMORY_OVERALLOCATION_BEHAVIOR_SPEC_VERSION                = 1
	KHR_PIPELINE_LIBRARY_EXTENSION_NAME                            = "VK_KHR_pipeline_library"
	EXT_PRIMITIVE_TOPOLOGY_LIST_RESTART_EXTENSION_NAME             = "VK_EXT_primitive_topology_list_restart"
	AMD_PIPELINE_COMPILER_CONTROL_SPEC_VERSION                     = 1
	EXT_YCBCR_IMAGE_ARRAYS_EXTENSION_NAME                          = "VK_EXT_ycbcr_image_arrays"
	KHR_SWAPCHAIN_EXTENSION_NAME                                   = "VK_KHR_swapchain"
	NVX_BINARY_IMPORT_SPEC_VERSION                                 = 1
	AMD_MEMORY_OVERALLOCATION_BEHAVIOR_EXTENSION_NAME              = "VK_AMD_memory_overallocation_behavior"
	KHR_MAINTENANCE1_SPEC_VERSION                                  = 2
	EXT_DEVICE_MEMORY_REPORT_EXTENSION_NAME                        = "VK_EXT_device_memory_report"
	KHR_COPY_COMMANDS_2_SPEC_VERSION                               = 1
	NV_EXTERNAL_MEMORY_RDMA_EXTENSION_NAME                         = "VK_NV_external_memory_rdma"
	KHR_SHADER_FLOAT16_INT8_EXTENSION_NAME                         = "VK_KHR_shader_float16_int8"
	NV_MESH_SHADER_SPEC_VERSION                                    = 1
	EXT_DEBUG_REPORT_SPEC_VERSION                                  = 10
	KHR_SHADER_TERMINATE_INVOCATION_SPEC_VERSION                   = 1
	KHR_EXTERNAL_MEMORY_FD_SPEC_VERSION                            = 1
	AMD_SHADER_INFO_SPEC_VERSION                                   = 1
	AMD_SHADER_FRAGMENT_MASK_EXTENSION_NAME                        = "VK_AMD_shader_fragment_mask"
	KHR_SEPARATE_DEPTH_STENCIL_LAYOUTS_EXTENSION_NAME              = "VK_KHR_separate_depth_stencil_layouts"
	NV_COOPERATIVE_MATRIX_EXTENSION_NAME                           = "VK_NV_cooperative_matrix"
	KHR_EXTERNAL_SEMAPHORE_FD_SPEC_VERSION                         = 1
	NV_DEVICE_GENERATED_COMMANDS_EXTENSION_NAME                    = "VK_NV_device_generated_commands"
	LUID_SIZE_KHR                                           uint32 = LUID_SIZE
	KHR_EXTERNAL_MEMORY_FD_EXTENSION_NAME                          = "VK_KHR_external_memory_fd"
	GOOGLE_DECORATE_STRING_EXTENSION_NAME                          = "VK_GOOGLE_decorate_string"
	KHR_SHADER_CLOCK_SPEC_VERSION                                  = 1
	NV_DEVICE_DIAGNOSTICS_CONFIG_EXTENSION_NAME                    = "VK_NV_device_diagnostics_config"
	IMG_FORMAT_PVRTC_SPEC_VERSION                                  = 1
	NV_REPRESENTATIVE_FRAGMENT_TEST_EXTENSION_NAME                 = "VK_NV_representative_fragment_test"
	EXT_QUEUE_FAMILY_FOREIGN_EXTENSION_NAME                        = "VK_EXT_queue_family_foreign"
	EXT_TEXTURE_COMPRESSION_ASTC_HDR_SPEC_VERSION                  = 1
	KHR_MAINTENANCE2_EXTENSION_NAME                                = "VK_KHR_maintenance2"
	EXT_ASTC_DECODE_MODE_EXTENSION_NAME                            = "VK_EXT_astc_decode_mode"
	EXT_DIRECT_MODE_DISPLAY_EXTENSION_NAME                         = "VK_EXT_direct_mode_display"
	AMD_RASTERIZATION_ORDER_SPEC_VERSION                           = 1
	NV_SHADER_IMAGE_FOOTPRINT_EXTENSION_NAME                       = "VK_NV_shader_image_footprint"
	KHR_PIPELINE_EXECUTABLE_PROPERTIES_EXTENSION_NAME              = "VK_KHR_pipeline_executable_properties"
	KHR_DEDICATED_ALLOCATION_EXTENSION_NAME                        = "VK_KHR_dedicated_allocation"
	KHR_DISPLAY_SPEC_VERSION                                       = 23
	EXT_YCBCR_IMAGE_ARRAYS_SPEC_VERSION                            = 1
	NV_SHADING_RATE_IMAGE_EXTENSION_NAME                           = "VK_NV_shading_rate_image"
	AMD_TEXTURE_GATHER_BIAS_LOD_SPEC_VERSION                       = 1
	EXT_VALIDATION_FEATURES_SPEC_VERSION                           = 5
	QCOM_RENDER_PASS_TRANSFORM_SPEC_VERSION                        = 2
	NV_COMPUTE_SHADER_DERIVATIVES_SPEC_VERSION                     = 1
	KHR_DEPTH_STENCIL_RESOLVE_EXTENSION_NAME                       = "VK_KHR_depth_stencil_resolve"
	KHR_EXTERNAL_MEMORY_CAPABILITIES_EXTENSION_NAME                = "VK_KHR_external_memory_capabilities"
	EXT_TEXTURE_COMPRESSION_ASTC_HDR_EXTENSION_NAME                = "VK_EXT_texture_compression_astc_hdr"
	IMG_FORMAT_PVRTC_EXTENSION_NAME                                = "VK_IMG_format_pvrtc"
	NV_DEDICATED_ALLOCATION_SPEC_VERSION                           = 1
	AMD_BUFFER_MARKER_SPEC_VERSION                                 = 1
	KHR_SHADER_FLOAT16_INT8_SPEC_VERSION                           = 1
	KHR_SHADER_INTEGER_DOT_PRODUCT_SPEC_VERSION                    = 1
	KHR_RAY_TRACING_PIPELINE_SPEC_VERSION                          = 1
	EXT_VERTEX_INPUT_DYNAMIC_STATE_EXTENSION_NAME                  = "VK_EXT_vertex_input_dynamic_state"
	EXT_DEBUG_REPORT_EXTENSION_NAME                                = "VK_EXT_debug_report"
	NV_SHADER_SM_BUILTINS_EXTENSION_NAME                           = "VK_NV_shader_sm_builtins"
	EXT_SAMPLE_LOCATIONS_SPEC_VERSION                              = 1
	NV_COMPUTE_SHADER_DERIVATIVES_EXTENSION_NAME                   = "VK_NV_compute_shader_derivatives"
	KHR_PUSH_DESCRIPTOR_SPEC_VERSION                               = 2
	NV_VIEWPORT_ARRAY2_SPEC_VERSION                                = 1
	NV_FRAGMENT_SHADER_BARYCENTRIC_EXTENSION_NAME                  = "VK_NV_fragment_shader_barycentric"
	NV_FRAMEBUFFER_MIXED_SAMPLES_EXTENSION_NAME                    = "VK_NV_framebuffer_mixed_samples"
	AMD_SHADER_IMAGE_LOAD_STORE_LOD_SPEC_VERSION                   = 1
	KHR_GET_PHYSICAL_DEVICE_PROPERTIES_2_SPEC_VERSION              = 2
	QCOM_RENDER_PASS_SHADER_RESOLVE_SPEC_VERSION                   = 4
	EXT_PRIVATE_DATA_EXTENSION_NAME                                = "VK_EXT_private_data"
	AMD_DRAW_INDIRECT_COUNT_EXTENSION_NAME                         = "VK_AMD_draw_indirect_count"
	EXT_FILTER_CUBIC_SPEC_VERSION                                  = 3
	KHR_BIND_MEMORY_2_SPEC_VERSION                                 = 1
	EXT_EXTENDED_DYNAMIC_STATE_EXTENSION_NAME                      = "VK_EXT_extended_dynamic_state"
	KHR_IMAGELESS_FRAMEBUFFER_EXTENSION_NAME                       = "VK_KHR_imageless_framebuffer"
	KHR_SHADER_TERMINATE_INVOCATION_EXTENSION_NAME                 = "VK_KHR_shader_terminate_invocation"
	KHR_PUSH_DESCRIPTOR_EXTENSION_NAME                             = "VK_KHR_push_descriptor"
	EXT_PIPELINE_CREATION_CACHE_CONTROL_EXTENSION_NAME             = "VK_EXT_pipeline_creation_cache_control"
	AMD_BUFFER_MARKER_EXTENSION_NAME                               = "VK_AMD_buffer_marker"
	AMD_GPU_SHADER_HALF_FLOAT_EXTENSION_NAME                       = "VK_AMD_gpu_shader_half_float"
	EXT_DESCRIPTOR_INDEXING_SPEC_VERSION                           = 2
	EXT_CONSERVATIVE_RASTERIZATION_EXTENSION_NAME                  = "VK_EXT_conservative_rasterization"
	EXT_SCALAR_BLOCK_LAYOUT_SPEC_VERSION                           = 1
	KHR_EXTERNAL_FENCE_EXTENSION_NAME                              = "VK_KHR_external_fence"
	EXT_PCI_BUS_INFO_EXTENSION_NAME                                = "VK_EXT_pci_bus_info"
	NV_RAY_TRACING_SPEC_VERSION                                    = 3
	KHR_SURFACE_PROTECTED_CAPABILITIES_EXTENSION_NAME              = "VK_KHR_surface_protected_capabilities"
	KHR_SHADER_FLOAT_CONTROLS_EXTENSION_NAME                       = "VK_KHR_shader_float_controls"
	KHR_UNIFORM_BUFFER_STANDARD_LAYOUT_SPEC_VERSION                = 1
	KHR_PIPELINE_LIBRARY_SPEC_VERSION                              = 1
)
