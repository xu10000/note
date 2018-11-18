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
let g:airline_theme='luna'
" Trigger configuration. Do not use <tab> if you use https://github.com/Valloric/YouCompleteMe.

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
" 搜索快捷键
map <silent> <Space>sd :Ack<Space>
map <Space>sb :/
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
imap <Space>ds <ESC>:q<CR>
" 格式化代码
map  <C-S-l>  :!go fmt<CR>
"  设置<C-u> <C-w>
nmap <C-w> :<ESC>daw
map <C-u> :<ESC>d0
filetype off                  " required
colorscheme gruvbox

call vundle#begin()

Plugin 'VundleVim/Vundle.vim'
Plugin 'ryanoasis/vim-devicons'
Plugin 'morhetz/gruvbox'
Plugin 'kien/ctrlp.vim'
Plugin 'scrooloose/nerdtree'
Plugin 'mileszs/ack.vim'
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
Plugin  'sidorares/node-vim-debugger'
call vundle#end()            " required
" 记住光标位置
au BufReadPost * if line("'\"") > 0|if line("'\"") <= line("$")|exe("norm '\"")|else|exe "norm $"|endif|endif

filetype plugin indent on    " required

