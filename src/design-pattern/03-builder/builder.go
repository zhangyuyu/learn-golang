package builder

const Success = "Success"

// resourcePoolConfig 资源池配置，不可导出
type resourcePoolConfig struct {
	// name 名字，不可导出，必填
	name string
	// maxTotal 最大连接数，不可导出，必填
	maxTotal uint16
	// maxIdle 最大空闲连接数，不可导出，必填
	maxIdle uint16
	// minIdle 最小空闲连接数，不可导出，不必填，但填写必须大于0
	minIdle uint16
}

// Builder 建造者接口
type Builder interface {
	// Build 创建对象
	Build() int16
}

// ResourcePoolBuild 资源池建造者，实现Builder接口
type ResourcePoolBuild struct {
	// resourcePoolConfig 嵌套资源池配置，组合的方式
	resourcePoolConfig
}

// Build 创建对象
func (build *ResourcePoolBuild) Build() string {
	// 这一串的判断可以保证builder里要求必须设置的属性，都能够设置
	if len(build.name) == 0 {
		return "ErrCodeNameRequired"
	}
	if build.maxTotal == 0 {
		return "ErrCodeMaxTotalNotSet"
	}
	if build.maxIdle == 0 {
		return "ErrCodeMaxIdleNotSet"
	}

	// 对设置后的对象，做整体的逻辑性验证
	if build.minIdle > build.maxIdle {
		return "ErrCodeMinIdleMoreThanMaxIdle"
	}
	return Success
}

// SetName 设置名字，并进行校验
func (build *ResourcePoolBuild) SetName(name string) string {
	if len(name) == 0 {
		return "ErrCodeNameRequired"
	}
	build.name = name
	return Success
}

// SetMaxTotal 设置最大连接数，并进行校验
func (build *ResourcePoolBuild) SetMaxTotal(maxTotal uint16) string {
	if maxTotal <= 0 {
		return "ErrCodeMaxTotalMustMoreThanZero"
	}
	build.maxTotal = maxTotal
	return Success
}

// SetMaxIdle 设置最大空闲连接数，并信息校验
func (build *ResourcePoolBuild) SetMaxIdle(maxIdle uint16) string {
	if maxIdle <= 0 {
		return "ErrCodeMaxIdleMustMoreThanZero"
	}
	build.maxIdle = maxIdle
	return Success
}

// SetMinIdle 设置最小空闲连接数，并进行校验
func (build *ResourcePoolBuild) SetMinIdle(minIdle uint16) string {
	if minIdle <= 0 {
		return "ErrCodeMinIdleMustMoreThanZero"
	}
	build.minIdle = minIdle
	return Success
}
