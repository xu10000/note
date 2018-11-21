syntax enable
set background=dark
set nocompatible              " be iMproved, required
set encoding=UTF-8
set backspace=2
set number
set splitbelow
set cursorcolumn
set cursorline
"  命令模式输入，tab在状态栏显示补全
set completeopt=longest,menu
set wildmenu

" set the runtime path to include Vundle and initialize
set rtp+=~/.vim/bundle/Vundle.vim
set rtp+=~/.fzf

set ts=4
let g:airline#extensions#tabline#enabled = 1
let g:airline#extensions#tabline#tab_nr_type = 1
"let Tlist_Use_Left_Window=1
let g:tagbar_left=1
let Tlist_Show_One_File=1 
let Tlist_Exit_OnlyWindow=1
"设置airline箭头
let g:airline_powerline_fonts = 1
" 设置airline颜色
let g:airline_theme='bubblegum'
" Trigger configuration. Do not use <tab> if you use https://github.com/Valloric/YouCompleteMe.
" 设置关灯效果
let g:limelight_conceal_ctermfg = 'gray'
let g:limelight_conceal_ctermfg = 240
" 注释配置
let g:NERDSpaceDelims = 1
let g:NERDCustomDelimiters = { 'c': { 'left': '','right': '' } }

nmap <Leader>b :bp<CR>
nmap <Leader>f :bn<CR>

" 查看buffers
nmap <Leader>l :ls<CR>

" 通过索引快速跳转
nmap <Leader>1 :1b<CR>
nmap <Leader>2 :2b<CR>
nmap <Leader>3 :3b<CR>
nmap <Leader>4 :4b<CR>
nmap <Leader>5 :5b<CR>
nmap <Leader>6 :6b<CR>
nmap <Leader>7 :7b<CR>
nmap <Leader>8 :8b<CR>
nmap <Leader>9 :9b<CR>
nmap <Leader>0 :10b<CR>
nmap <F3> :NERDTreeToggle<CR>
nmap <F2> :TagbarToggle<CR>
"保存文件
map <silent> <c-s> :update<CR> 
" 进入模式
nmap <Space>c :
" 关闭当前buffer
nmap <Space>bd :bdelete<CR>
" 打开终端
nmap <Space>' :term<CR>
 "设置tab键映射
nmap <tab> :bn<cr>
" debug go
nmap <Space>dl :DlvDebug<CR>y<CR>:$<CR>i
nmap <Space>db :DlvToggleBreakpoint<CR>
nmap <Space>ds <ESC>:q<CR>
" 格式化代码
map  <C-S-l>  :!go fmt<CR>
"  设置<C-u>
nmap <C-u> :<ESC>d0
"  搜索
nmap  <Space>sf  :Files<CR>
nmap  <Space>sb  :Lines<CR>
nmap  <Space>sd  :Ag<Space>
nmap <Leader>l  :Limelight<CR>

filetype off                  " required
colorscheme gruvbox

call vundle#begin()

Plugin 'VundleVim/Vundle.vim'
Plugin 'ryanoasis/vim-devicons'
Plugin 'morhetz/gruvbox'
Plugin 'scrooloose/nerdtree'
Plugin 'vim-airline/vim-airline'
Plugin 'vim-airline/vim-airline-themes'
Plugin 'majutsushi/tagbar'
Plugin 'fatih/vim-go'
Plugin 'Valloric/YouCompleteMe'
" Track the engine.
Plugin 'SirVer/ultisnips'
" Snippets are separated from the engine. Add this if you want them:
Plugin 'honza/vim-snippets'
" debug go
Plugin 'Shougo/vimproc.vim', {'do' : 'make'}
Plugin 'Shougo/vimshell'
Plugin 'sebdah/vim-delve'
" 搜索引擎
Plugin 'junegunn/fzf.vim'
"  设置关灯效果
Plugin 'junegunn/limelight.vim'
" go 代码提示
Plugin 'nsf/gocode', {'rtp': 'vim/'}
" 注释插件
Plugin 'scrooloose/nerdcommenter'

call vundle#end()            " required
" 记住光标位置
au BufReadPost * if line("'\"") > 0|if line("'\"") <= line("$")|exe("norm '\"")|else|exe "norm $"|endif|endif
filetype plugin indent on    " required

