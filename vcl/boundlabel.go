
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "github.com/ying32/govcl/vcl/api"
    . "github.com/ying32/govcl/vcl/types"
    "unsafe"
)

type TBoundLabel struct {
    IControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewBoundLabel
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewBoundLabel(owner IComponent) *TBoundLabel {
    b := new(TBoundLabel)
    b.instance = BoundLabel_Create(CheckPtr(owner))
    b.ptr = unsafe.Pointer(b.instance)
    return b
}

// BoundLabelFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func BoundLabelFromInst(inst uintptr) *TBoundLabel {
    b := new(TBoundLabel)
    b.instance = inst
    b.ptr = unsafe.Pointer(inst)
    return b
}

// BoundLabelFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func BoundLabelFromObj(obj IObject) *TBoundLabel {
    b := new(TBoundLabel)
    b.instance = CheckPtr(obj)
    b.ptr = unsafe.Pointer(b.instance)
    return b
}

// BoundLabelFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func BoundLabelFromUnsafePointer(ptr unsafe.Pointer) *TBoundLabel {
    b := new(TBoundLabel)
    b.instance = uintptr(ptr)
    b.ptr = ptr
    return b
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (b *TBoundLabel) Free() {
    if b.instance != 0 {
        BoundLabel_Free(b.instance)
        b.instance = 0
        b.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (b *TBoundLabel) Instance() uintptr {
    return b.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (b *TBoundLabel) UnsafeAddr() unsafe.Pointer {
    return b.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (b *TBoundLabel) IsValid() bool {
    return b.instance != 0
}

// TBoundLabelClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TBoundLabelClass() TClass {
    return BoundLabel_StaticClassType()
}

// BringToFront
// CN: 将控件置于最前。
// EN: Bring the control to the front.
func (b *TBoundLabel) BringToFront() {
    BoundLabel_BringToFront(b.instance)
}

// ClientToScreen
// CN: 将客户端坐标转为绝对的屏幕坐标。
// EN: Convert client coordinates to absolute screen coordinates.
func (b *TBoundLabel) ClientToScreen(Point TPoint) TPoint {
    return BoundLabel_ClientToScreen(b.instance, Point)
}

// ClientToParent
// CN: 将客户端坐标转为父容器坐标。
// EN: Convert client coordinates to parent container coordinates.
func (b *TBoundLabel) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return BoundLabel_ClientToParent(b.instance, Point , CheckPtr(AParent))
}

// Dragging
// CN: 是否在拖拽中。
// EN: Is it in the middle of dragging.
func (b *TBoundLabel) Dragging() bool {
    return BoundLabel_Dragging(b.instance)
}

// HasParent
// CN: 是否有父容器。
// EN: Is there a parent container.
func (b *TBoundLabel) HasParent() bool {
    return BoundLabel_HasParent(b.instance)
}

// Hide
// CN: 隐藏控件。
// EN: Hidden control.
func (b *TBoundLabel) Hide() {
    BoundLabel_Hide(b.instance)
}

// Invalidate
// CN: 要求重绘。
// EN: Redraw.
func (b *TBoundLabel) Invalidate() {
    BoundLabel_Invalidate(b.instance)
}

// Perform
// CN: 发送一个消息。
// EN: Send a message.
func (b *TBoundLabel) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return BoundLabel_Perform(b.instance, Msg , WParam , LParam)
}

// Refresh
// CN: 刷新控件。
// EN: Refresh control.
func (b *TBoundLabel) Refresh() {
    BoundLabel_Refresh(b.instance)
}

// Repaint
// CN: 重绘。
// EN: Repaint.
func (b *TBoundLabel) Repaint() {
    BoundLabel_Repaint(b.instance)
}

// ScreenToClient
// CN: 将屏幕坐标转为客户端坐标。
// EN: Convert screen coordinates to client coordinates.
func (b *TBoundLabel) ScreenToClient(Point TPoint) TPoint {
    return BoundLabel_ScreenToClient(b.instance, Point)
}

// ParentToClient
// CN: 将父容器坐标转为客户端坐标。
// EN: Convert parent container coordinates to client coordinates.
func (b *TBoundLabel) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return BoundLabel_ParentToClient(b.instance, Point , CheckPtr(AParent))
}

// SendToBack
// CN: 控件至于最后面。
// EN: The control is placed at the end.
func (b *TBoundLabel) SendToBack() {
    BoundLabel_SendToBack(b.instance)
}

// SetBounds
// CN: 设置组件边界。
// EN: Set component boundaries.
func (b *TBoundLabel) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    BoundLabel_SetBounds(b.instance, ALeft , ATop , AWidth , AHeight)
}

// Show
// CN: 显示控件。
// EN: Show control.
func (b *TBoundLabel) Show() {
    BoundLabel_Show(b.instance)
}

// Update
// CN: 控件更新。
// EN: Update.
func (b *TBoundLabel) Update() {
    BoundLabel_Update(b.instance)
}

// GetTextBuf
// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (b *TBoundLabel) GetTextBuf(Buffer string, BufSize int32) int32 {
    return BoundLabel_GetTextBuf(b.instance, Buffer , BufSize)
}

// GetTextLen
// CN: 获取控件的字符长，如果有。
// EN: Get the character length of the control, if any.
func (b *TBoundLabel) GetTextLen() int32 {
    return BoundLabel_GetTextLen(b.instance)
}

// SetTextBuf
// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (b *TBoundLabel) SetTextBuf(Buffer string) {
    BoundLabel_SetTextBuf(b.instance, Buffer)
}

// FindComponent
// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (b *TBoundLabel) FindComponent(AName string) *TComponent {
    return ComponentFromInst(BoundLabel_FindComponent(b.instance, AName))
}

// GetNamePath
// CN: 获取类名路径。
// EN: Get the class name path.
func (b *TBoundLabel) GetNamePath() string {
    return BoundLabel_GetNamePath(b.instance)
}

// Assign
// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (b *TBoundLabel) Assign(Source IObject) {
    BoundLabel_Assign(b.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (b *TBoundLabel) DisposeOf() {
    BoundLabel_DisposeOf(b.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (b *TBoundLabel) ClassType() TClass {
    return BoundLabel_ClassType(b.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (b *TBoundLabel) ClassName() string {
    return BoundLabel_ClassName(b.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (b *TBoundLabel) InstanceSize() int32 {
    return BoundLabel_InstanceSize(b.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (b *TBoundLabel) InheritsFrom(AClass TClass) bool {
    return BoundLabel_InheritsFrom(b.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (b *TBoundLabel) Equals(Obj IObject) bool {
    return BoundLabel_Equals(b.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (b *TBoundLabel) GetHashCode() int32 {
    return BoundLabel_GetHashCode(b.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (b *TBoundLabel) ToString() string {
    return BoundLabel_ToString(b.instance)
}

// BiDiMode
func (b *TBoundLabel) BiDiMode() TBiDiMode {
    return BoundLabel_GetBiDiMode(b.instance)
}

// SetBiDiMode
func (b *TBoundLabel) SetBiDiMode(value TBiDiMode) {
    BoundLabel_SetBiDiMode(b.instance, value)
}

// Caption
// CN: 获取控件标题。
// EN: Get the control title.
func (b *TBoundLabel) Caption() string {
    return BoundLabel_GetCaption(b.instance)
}

// SetCaption
// CN: 设置控件标题。
// EN: Set the control title.
func (b *TBoundLabel) SetCaption(value string) {
    BoundLabel_SetCaption(b.instance, value)
}

// Color
// CN: 获取颜色。
// EN: Get color.
func (b *TBoundLabel) Color() TColor {
    return BoundLabel_GetColor(b.instance)
}

// SetColor
// CN: 设置颜色。
// EN: Set color.
func (b *TBoundLabel) SetColor(value TColor) {
    BoundLabel_SetColor(b.instance, value)
}

// DragCursor
// CN: 获取设置控件拖拽时的光标。
// EN: Get Set the cursor when the control is dragged.
func (b *TBoundLabel) DragCursor() TCursor {
    return BoundLabel_GetDragCursor(b.instance)
}

// SetDragCursor
// CN: 设置设置控件拖拽时的光标。
// EN: Set Set the cursor when the control is dragged.
func (b *TBoundLabel) SetDragCursor(value TCursor) {
    BoundLabel_SetDragCursor(b.instance, value)
}

// DragKind
// CN: 获取拖拽方式。
// EN: Get Drag and drop.
func (b *TBoundLabel) DragKind() TDragKind {
    return BoundLabel_GetDragKind(b.instance)
}

// SetDragKind
// CN: 设置拖拽方式。
// EN: Set Drag and drop.
func (b *TBoundLabel) SetDragKind(value TDragKind) {
    BoundLabel_SetDragKind(b.instance, value)
}

// DragMode
// CN: 获取拖拽模式。
// EN: Get Drag mode.
func (b *TBoundLabel) DragMode() TDragMode {
    return BoundLabel_GetDragMode(b.instance)
}

// SetDragMode
// CN: 设置拖拽模式。
// EN: Set Drag mode.
func (b *TBoundLabel) SetDragMode(value TDragMode) {
    BoundLabel_SetDragMode(b.instance, value)
}

// Font
// CN: 获取字体。
// EN: Get Font.
func (b *TBoundLabel) Font() *TFont {
    return FontFromInst(BoundLabel_GetFont(b.instance))
}

// SetFont
// CN: 设置字体。
// EN: Set Font.
func (b *TBoundLabel) SetFont(value *TFont) {
    BoundLabel_SetFont(b.instance, CheckPtr(value))
}

// Height
// CN: 获取高度。
// EN: Get height.
func (b *TBoundLabel) Height() int32 {
    return BoundLabel_GetHeight(b.instance)
}

// SetHeight
// CN: 设置高度。
// EN: Set height.
func (b *TBoundLabel) SetHeight(value int32) {
    BoundLabel_SetHeight(b.instance, value)
}

// Left
// CN: 获取左边位置。
// EN: Get Left position.
func (b *TBoundLabel) Left() int32 {
    return BoundLabel_GetLeft(b.instance)
}

// ParentColor
// CN: 获取父容器颜色。
// EN: Get parent color.
func (b *TBoundLabel) ParentColor() bool {
    return BoundLabel_GetParentColor(b.instance)
}

// SetParentColor
// CN: 设置父容器颜色。
// EN: Set parent color.
func (b *TBoundLabel) SetParentColor(value bool) {
    BoundLabel_SetParentColor(b.instance, value)
}

// ParentFont
// CN: 获取父容器字体。
// EN: Get Parent container font.
func (b *TBoundLabel) ParentFont() bool {
    return BoundLabel_GetParentFont(b.instance)
}

// SetParentFont
// CN: 设置父容器字体。
// EN: Set Parent container font.
func (b *TBoundLabel) SetParentFont(value bool) {
    BoundLabel_SetParentFont(b.instance, value)
}

// ParentShowHint
func (b *TBoundLabel) ParentShowHint() bool {
    return BoundLabel_GetParentShowHint(b.instance)
}

// SetParentShowHint
func (b *TBoundLabel) SetParentShowHint(value bool) {
    BoundLabel_SetParentShowHint(b.instance, value)
}

// PopupMenu
// CN: 获取右键菜单。
// EN: Get Right click menu.
func (b *TBoundLabel) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(BoundLabel_GetPopupMenu(b.instance))
}

// SetPopupMenu
// CN: 设置右键菜单。
// EN: Set Right click menu.
func (b *TBoundLabel) SetPopupMenu(value IComponent) {
    BoundLabel_SetPopupMenu(b.instance, CheckPtr(value))
}

// ShowAccelChar
func (b *TBoundLabel) ShowAccelChar() bool {
    return BoundLabel_GetShowAccelChar(b.instance)
}

// SetShowAccelChar
func (b *TBoundLabel) SetShowAccelChar(value bool) {
    BoundLabel_SetShowAccelChar(b.instance, value)
}

// ShowHint
// CN: 获取显示鼠标悬停提示。
// EN: Get Show mouseover tips.
func (b *TBoundLabel) ShowHint() bool {
    return BoundLabel_GetShowHint(b.instance)
}

// SetShowHint
// CN: 设置显示鼠标悬停提示。
// EN: Set Show mouseover tips.
func (b *TBoundLabel) SetShowHint(value bool) {
    BoundLabel_SetShowHint(b.instance, value)
}

// Top
// CN: 获取顶边位置。
// EN: Get Top position.
func (b *TBoundLabel) Top() int32 {
    return BoundLabel_GetTop(b.instance)
}

// Transparent
// CN: 获取透明。
// EN: Get transparent.
func (b *TBoundLabel) Transparent() bool {
    return BoundLabel_GetTransparent(b.instance)
}

// SetTransparent
// CN: 设置透明。
// EN: Set transparent.
func (b *TBoundLabel) SetTransparent(value bool) {
    BoundLabel_SetTransparent(b.instance, value)
}

// Layout
func (b *TBoundLabel) Layout() TTextLayout {
    return BoundLabel_GetLayout(b.instance)
}

// SetLayout
func (b *TBoundLabel) SetLayout(value TTextLayout) {
    BoundLabel_SetLayout(b.instance, value)
}

// WordWrap
// CN: 获取自动换行。
// EN: Get Automatic line break.
func (b *TBoundLabel) WordWrap() bool {
    return BoundLabel_GetWordWrap(b.instance)
}

// SetWordWrap
// CN: 设置自动换行。
// EN: Set Automatic line break.
func (b *TBoundLabel) SetWordWrap(value bool) {
    BoundLabel_SetWordWrap(b.instance, value)
}

// Width
// CN: 获取宽度。
// EN: Get width.
func (b *TBoundLabel) Width() int32 {
    return BoundLabel_GetWidth(b.instance)
}

// SetWidth
// CN: 设置宽度。
// EN: Set width.
func (b *TBoundLabel) SetWidth(value int32) {
    BoundLabel_SetWidth(b.instance, value)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (b *TBoundLabel) SetOnClick(fn TNotifyEvent) {
    BoundLabel_SetOnClick(b.instance, fn)
}

// SetOnContextPopup
// CN: 设置上下文弹出事件，一般是右键时弹出。
// EN: Set Context popup event, usually pop up when right click.
func (b *TBoundLabel) SetOnContextPopup(fn TContextPopupEvent) {
    BoundLabel_SetOnContextPopup(b.instance, fn)
}

// SetOnDblClick
// CN: 设置双击事件。
// EN: .
func (b *TBoundLabel) SetOnDblClick(fn TNotifyEvent) {
    BoundLabel_SetOnDblClick(b.instance, fn)
}

// SetOnDragDrop
// CN: 设置拖拽下落事件。
// EN: Set Drag and drop event.
func (b *TBoundLabel) SetOnDragDrop(fn TDragDropEvent) {
    BoundLabel_SetOnDragDrop(b.instance, fn)
}

// SetOnDragOver
// CN: 设置拖拽完成事件。
// EN: Set Drag and drop completion event.
func (b *TBoundLabel) SetOnDragOver(fn TDragOverEvent) {
    BoundLabel_SetOnDragOver(b.instance, fn)
}

// SetOnEndDock
// CN: 设置停靠结束事件。
// EN: Set Dock end event.
func (b *TBoundLabel) SetOnEndDock(fn TEndDragEvent) {
    BoundLabel_SetOnEndDock(b.instance, fn)
}

// SetOnEndDrag
// CN: 设置拖拽结束。
// EN: Set End of drag.
func (b *TBoundLabel) SetOnEndDrag(fn TEndDragEvent) {
    BoundLabel_SetOnEndDrag(b.instance, fn)
}

// SetOnGesture
func (b *TBoundLabel) SetOnGesture(fn TGestureEvent) {
    BoundLabel_SetOnGesture(b.instance, fn)
}

// SetOnMouseActivate
func (b *TBoundLabel) SetOnMouseActivate(fn TMouseActivateEvent) {
    BoundLabel_SetOnMouseActivate(b.instance, fn)
}

// SetOnMouseDown
// CN: 设置鼠标按下事件。
// EN: Set Mouse down event.
func (b *TBoundLabel) SetOnMouseDown(fn TMouseEvent) {
    BoundLabel_SetOnMouseDown(b.instance, fn)
}

// SetOnMouseMove
// CN: 设置鼠标移动事件。
// EN: .
func (b *TBoundLabel) SetOnMouseMove(fn TMouseMoveEvent) {
    BoundLabel_SetOnMouseMove(b.instance, fn)
}

// SetOnMouseUp
// CN: 设置鼠标抬起事件。
// EN: Set Mouse lift event.
func (b *TBoundLabel) SetOnMouseUp(fn TMouseEvent) {
    BoundLabel_SetOnMouseUp(b.instance, fn)
}

// SetOnStartDock
// CN: 设置启动停靠。
// EN: .
func (b *TBoundLabel) SetOnStartDock(fn TStartDockEvent) {
    BoundLabel_SetOnStartDock(b.instance, fn)
}

// Canvas
// CN: 获取画布。
// EN: .
func (b *TBoundLabel) Canvas() *TCanvas {
    return CanvasFromInst(BoundLabel_GetCanvas(b.instance))
}

// GlowSize
func (b *TBoundLabel) GlowSize() int32 {
    return BoundLabel_GetGlowSize(b.instance)
}

// SetGlowSize
func (b *TBoundLabel) SetGlowSize(value int32) {
    BoundLabel_SetGlowSize(b.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (b *TBoundLabel) Enabled() bool {
    return BoundLabel_GetEnabled(b.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (b *TBoundLabel) SetEnabled(value bool) {
    BoundLabel_SetEnabled(b.instance, value)
}

// Action
func (b *TBoundLabel) Action() *TAction {
    return ActionFromInst(BoundLabel_GetAction(b.instance))
}

// SetAction
func (b *TBoundLabel) SetAction(value IComponent) {
    BoundLabel_SetAction(b.instance, CheckPtr(value))
}

// Align
// CN: 获取控件自动调整。
// EN: Get Control automatically adjusts.
func (b *TBoundLabel) Align() TAlign {
    return BoundLabel_GetAlign(b.instance)
}

// SetAlign
// CN: 设置控件自动调整。
// EN: Set Control automatically adjusts.
func (b *TBoundLabel) SetAlign(value TAlign) {
    BoundLabel_SetAlign(b.instance, value)
}

// Anchors
// CN: 获取四个角位置的锚点。
// EN: .
func (b *TBoundLabel) Anchors() TAnchors {
    return BoundLabel_GetAnchors(b.instance)
}

// SetAnchors
// CN: 设置四个角位置的锚点。
// EN: .
func (b *TBoundLabel) SetAnchors(value TAnchors) {
    BoundLabel_SetAnchors(b.instance, value)
}

// BoundsRect
func (b *TBoundLabel) BoundsRect() TRect {
    return BoundLabel_GetBoundsRect(b.instance)
}

// SetBoundsRect
func (b *TBoundLabel) SetBoundsRect(value TRect) {
    BoundLabel_SetBoundsRect(b.instance, value)
}

// ClientHeight
// CN: 获取客户区高度。
// EN: Get client height.
func (b *TBoundLabel) ClientHeight() int32 {
    return BoundLabel_GetClientHeight(b.instance)
}

// SetClientHeight
// CN: 设置客户区高度。
// EN: Set client height.
func (b *TBoundLabel) SetClientHeight(value int32) {
    BoundLabel_SetClientHeight(b.instance, value)
}

// ClientOrigin
func (b *TBoundLabel) ClientOrigin() TPoint {
    return BoundLabel_GetClientOrigin(b.instance)
}

// ClientRect
// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (b *TBoundLabel) ClientRect() TRect {
    return BoundLabel_GetClientRect(b.instance)
}

// ClientWidth
// CN: 获取客户区宽度。
// EN: Get client width.
func (b *TBoundLabel) ClientWidth() int32 {
    return BoundLabel_GetClientWidth(b.instance)
}

// SetClientWidth
// CN: 设置客户区宽度。
// EN: Set client width.
func (b *TBoundLabel) SetClientWidth(value int32) {
    BoundLabel_SetClientWidth(b.instance, value)
}

// ControlState
// CN: 获取控件状态。
// EN: Get control state.
func (b *TBoundLabel) ControlState() TControlState {
    return BoundLabel_GetControlState(b.instance)
}

// SetControlState
// CN: 设置控件状态。
// EN: Set control state.
func (b *TBoundLabel) SetControlState(value TControlState) {
    BoundLabel_SetControlState(b.instance, value)
}

// ControlStyle
// CN: 获取控件样式。
// EN: Get control style.
func (b *TBoundLabel) ControlStyle() TControlStyle {
    return BoundLabel_GetControlStyle(b.instance)
}

// SetControlStyle
// CN: 设置控件样式。
// EN: Set control style.
func (b *TBoundLabel) SetControlStyle(value TControlStyle) {
    BoundLabel_SetControlStyle(b.instance, value)
}

// ExplicitLeft
func (b *TBoundLabel) ExplicitLeft() int32 {
    return BoundLabel_GetExplicitLeft(b.instance)
}

// ExplicitTop
func (b *TBoundLabel) ExplicitTop() int32 {
    return BoundLabel_GetExplicitTop(b.instance)
}

// ExplicitWidth
func (b *TBoundLabel) ExplicitWidth() int32 {
    return BoundLabel_GetExplicitWidth(b.instance)
}

// ExplicitHeight
func (b *TBoundLabel) ExplicitHeight() int32 {
    return BoundLabel_GetExplicitHeight(b.instance)
}

// Floating
func (b *TBoundLabel) Floating() bool {
    return BoundLabel_GetFloating(b.instance)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (b *TBoundLabel) Visible() bool {
    return BoundLabel_GetVisible(b.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (b *TBoundLabel) SetVisible(value bool) {
    BoundLabel_SetVisible(b.instance, value)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (b *TBoundLabel) Parent() *TWinControl {
    return WinControlFromInst(BoundLabel_GetParent(b.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (b *TBoundLabel) SetParent(value IWinControl) {
    BoundLabel_SetParent(b.instance, CheckPtr(value))
}

// StyleElements
// CN: 获取样式元素。
// EN: Get Style element.
func (b *TBoundLabel) StyleElements() TStyleElements {
    return BoundLabel_GetStyleElements(b.instance)
}

// SetStyleElements
// CN: 设置样式元素。
// EN: Set Style element.
func (b *TBoundLabel) SetStyleElements(value TStyleElements) {
    BoundLabel_SetStyleElements(b.instance, value)
}

// AlignWithMargins
// CN: 获取边距，仅VCL有效。
// EN: Get The margin. Only VCL is valid..
func (b *TBoundLabel) AlignWithMargins() bool {
    return BoundLabel_GetAlignWithMargins(b.instance)
}

// SetAlignWithMargins
// CN: 设置边距，仅VCL有效。
// EN: Set The margin. Only VCL is valid..
func (b *TBoundLabel) SetAlignWithMargins(value bool) {
    BoundLabel_SetAlignWithMargins(b.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (b *TBoundLabel) Cursor() TCursor {
    return BoundLabel_GetCursor(b.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (b *TBoundLabel) SetCursor(value TCursor) {
    BoundLabel_SetCursor(b.instance, value)
}

// Hint
// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (b *TBoundLabel) Hint() string {
    return BoundLabel_GetHint(b.instance)
}

// SetHint
// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (b *TBoundLabel) SetHint(value string) {
    BoundLabel_SetHint(b.instance, value)
}

// Margins
// CN: 获取边矩，仅VCL有效。
// EN: Get Edge moment, only VCL is valid.
func (b *TBoundLabel) Margins() *TMargins {
    return MarginsFromInst(BoundLabel_GetMargins(b.instance))
}

// SetMargins
// CN: 设置边矩，仅VCL有效。
// EN: Set Edge moment, only VCL is valid.
func (b *TBoundLabel) SetMargins(value *TMargins) {
    BoundLabel_SetMargins(b.instance, CheckPtr(value))
}

// CustomHint
// CN: 获取自定义提示。
// EN: Get custom hint.
func (b *TBoundLabel) CustomHint() *TCustomHint {
    return CustomHintFromInst(BoundLabel_GetCustomHint(b.instance))
}

// SetCustomHint
// CN: 设置自定义提示。
// EN: Set custom hint.
func (b *TBoundLabel) SetCustomHint(value IComponent) {
    BoundLabel_SetCustomHint(b.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (b *TBoundLabel) ComponentCount() int32 {
    return BoundLabel_GetComponentCount(b.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (b *TBoundLabel) ComponentIndex() int32 {
    return BoundLabel_GetComponentIndex(b.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (b *TBoundLabel) SetComponentIndex(value int32) {
    BoundLabel_SetComponentIndex(b.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (b *TBoundLabel) Owner() *TComponent {
    return ComponentFromInst(BoundLabel_GetOwner(b.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (b *TBoundLabel) Name() string {
    return BoundLabel_GetName(b.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (b *TBoundLabel) SetName(value string) {
    BoundLabel_SetName(b.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (b *TBoundLabel) Tag() int {
    return BoundLabel_GetTag(b.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (b *TBoundLabel) SetTag(value int) {
    BoundLabel_SetTag(b.instance, value)
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (b *TBoundLabel) Components(AIndex int32) *TComponent {
    return ComponentFromInst(BoundLabel_GetComponents(b.instance, AIndex))
}
